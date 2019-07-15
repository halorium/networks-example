package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/state"
)

func getNetworkAccountPurchases(state *state.State) {
	logger.Debug("get-network-account-purchases", state)

	networkAccountPurchaseEntity := networkaccountpurchase.NewEntity()
	networkAccountPurchaseEntity.Head.NetworkID = state.Request.Params.ByName("network-id")
	networkAccountPurchaseEntity.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")

	collection, flawError := entities.GetEntitiesFromPostgres(networkAccountPurchaseEntity)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworkAccountPurchases")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = nil
	state.Response.Body = collection
	state.Response.Ok()
}
