package main

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/affiliate-window/transaction-loader"
	"github.com/halorium/networks-example/networks/affilinet/transaction-loader"
	"github.com/halorium/networks-example/networks/booking/transaction-loader"
	"github.com/halorium/networks-example/networks/rakuten/transaction-loader"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/state"
)

func postNetworkAccountPurchasesLoad(message *notifier.Message) *flaw.Error {
	state, flawError := state.NewFromMessage(message)

	if flawError != nil {
		return flawError.Wrap("cannot postNetworkAccountPurchasesLoad")
	}

	logger.Debug("worker-post-network-account-purchases-load", state)

	switch state.Network.Body.Variant {
	case affiliatewindow.Variant:
		return affiliatewindow.LoadTransactions(state)
	case affilinet.Variant:
		return affilinet.LoadTransactions(state)
	case booking.Variant:
		return booking.LoadTransactions(state)
	case rakuten.Variant:
		return rakuten.LoadTransactions(state)
	}

	return flaw.New("attempt to load invalid variant: " + state.Network.Body.Variant).Wrap("cannot postNetworkAccountPurchasesLoad")
}
