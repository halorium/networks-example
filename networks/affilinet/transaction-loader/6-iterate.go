package affilinet

import (
	"fmt"

	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func iterate(state *state.State) *flaw.Error {
	logger.Debug("affilinet-iterate", state)

	res, ok := state.VariantResponse.(response)

	if !ok {
		return flaw.New(fmt.Sprintf("invalid type assertion: %T", res)).Wrap("cannot iterate")
	}

	failedTransactionCount := 0

	var flawError *flaw.Error

	for _, transaction := range res.Transactions {
		logger.Debug("affilinet-iterate-transaction", transaction)

		// this is not a critical error, not worth passing up...
		_ = state.Message.Refresh()

		state.NetworkAccountPurchase, flawError = newPurchaseFromTransaction(state, transaction)

		if flawError != nil {
			logger.Warn("affilinet-iterate-transaction-error", flawError.Wrap("cannot iterate"))
			failedTransactionCount++
			continue
		}

		flawError = entities.PutToAPI(envvars.NetworksHost, state.NetworkAccountPurchase)

		if flawError != nil {
			logger.Warn("affilinet-iterate-transaction-error", flawError.Wrap("cannot iterate"))
			failedTransactionCount++
		}
	}

	if failedTransactionCount > 0 {
		return flaw.New(fmt.Sprintf("cannot iterate: %v transactions failed", failedTransactionCount))
	}

	return nil
}
