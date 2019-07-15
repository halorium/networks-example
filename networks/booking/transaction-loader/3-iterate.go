package booking

import (
	"fmt"

	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func iterate(state *state.State) *flaw.Error {
	logger.Debug("booking-iterate", state)

	transactions, ok := state.VariantResponse.([]transaction)

	if !ok {
		return flaw.New(fmt.Sprintf("invalid type assertion: %T", transactions)).Wrap("cannot iterate")
	}

	failedTransactionCount := 0

	var flawError *flaw.Error

	for _, transaction := range transactions {
		logger.Debug("booking-iterate-transaction", transaction)

		// this is not a critical error, not worth passing up...
		_ = state.Message.Refresh()

		if state.NetworkAccount.Body.Variations.AdditionalID != transaction.AffiliateID {
			// skip records with invalid affiliate-id as per https://we-like-it.atlassian.net/browse/INTICKETS-35
			logger.Warn("booking-iterate-transaction-invalid-affiliate-id", transaction)
			continue
		}

		state.NetworkAccountPurchase, flawError = newPurchaseFromTransaction(state, transaction)

		if flawError != nil {
			logger.Warn("booking-iterate-transaction-error", flawError.Wrap("cannot iterate"))
			failedTransactionCount++
			continue
		}

		flawError = entities.PutToAPI(envvars.NetworksHost, state.NetworkAccountPurchase)

		if flawError != nil {
			logger.Warn("booking-iterate-transaction-error", flawError.Wrap("cannot iterate"))
			failedTransactionCount++
		}
	}

	if failedTransactionCount > 0 {
		return flaw.New(fmt.Sprintf("cannot iterate: %v transactions failed", failedTransactionCount))
	}

	return nil
}
