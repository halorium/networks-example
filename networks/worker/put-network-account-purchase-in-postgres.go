package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/state"
)

func putNetworkAccountPurchaseInPostgres(message *notifier.Message) *flaw.Error {
	state, flawError := state.NewFromMessage(message)

	if flawError != nil {
		return flawError.Wrap("cannot putNetworkAccountPurchaseInPostgres")
	}

	logger.Debug("worker-put-network-account-purchase-in-postgres", state)

	flawError = entities.PutEntityInPostgres(state.NetworkAccountPurchase)

	if flawError != nil {
		return flawError.Wrap("cannot putNetworkAccountPurchaseInPostgres")
	}

	return nil
}
