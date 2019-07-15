package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/state"
)

func getNetworkAccountPurchase(state *state.State) {
	logger.Debug("get-network-account-purchase", state)

	state.NetworkAccountPurchase = networkaccountpurchase.NewEntity()

	state.NetworkAccountPurchase.Head.NetworkID = state.Request.Params.ByName("network-id")
	state.NetworkAccountPurchase.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")
	state.NetworkAccountPurchase.Head.PurchaseID = state.Request.Params.ByName("purchase-id")

	flawError := entities.GetEntityFromPostgres(state.NetworkAccountPurchase)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworkAccountPurchase")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.NetworkAccountPurchase.Head
	state.Response.Body = state.NetworkAccountPurchase.Body
	state.Response.Ok()
}
