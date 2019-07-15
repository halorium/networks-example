package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func deleteNetworkAccountPurchases(state *state.State) {
	logger.Debug("delete-network-account-purchases", state)

	state.NetworkAccountPurchase = networkaccountpurchase.NewEntity()
	state.NetworkAccountPurchase.Head.NetworkID = state.Request.Params.ByName("network-id")
	state.NetworkAccountPurchase.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")

	flawError := entities.DeleteEntitiesFromPostgres(state.NetworkAccountPurchase)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot deleteNetworkAccountPurchases")

		if flawError.Error() == "subordinate entities exist" {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("subordinate entities exist")
			state.Response.Conflict()
			return
		}

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Ok()
}
