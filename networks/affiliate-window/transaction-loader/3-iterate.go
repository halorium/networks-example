package affiliatewindow

import (
	"fmt"

	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
)

func iterate(state *state.State) *flaw.Error {
	logger.Debug("affiliate-window-iterate", state)

	transactions, ok := state.VariantResponse.([]rawTransaction)

	if !ok {
		return flaw.New(fmt.Sprintf("invalid type assertion: %T", transactions)).Wrap("cannot iterate")
	}

	failedTransactionCount := 0

	var flawError *flaw.Error

	for _, original := range transactions {
		logger.Debug("affiliate-window-iterate-transaction", original)

		// this is not a critical error, not worth passing up...
		_ = state.Message.Refresh()

		newTransaction := transaction{}

		serializers.CopyAndStringify(&original, &newTransaction)

		state.NetworkAccountPurchase, flawError = newPurchaseFromTransaction(state, newTransaction)

		if flawError != nil {
			logger.Warn("affiliate-window-iterate-transaction-error", flawError.Wrap("cannot iterate"))
			failedTransactionCount++
			continue
		}

		flawError = entities.PutToAPI(envvars.NetworksHost, state.NetworkAccountPurchase)

		if flawError != nil {
			logger.Warn("affiliate-window-iterate-transaction-error", flawError.Wrap("cannot iterate"))
			failedTransactionCount++
		}
	}

	if failedTransactionCount > 0 {
		return flaw.New(fmt.Sprintf("cannot iterate: %v transactions failed", failedTransactionCount))
	}

	return nil
}
