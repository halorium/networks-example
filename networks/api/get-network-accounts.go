package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account"
	"github.com/halorium/networks-example/state"
)

func getNetworkAccounts(state *state.State) {
	logger.Debug("get-network-accounts", state)

	networkAccountEntity := networkaccount.NewEntity()
	networkAccountEntity.Head.NetworkID = state.Request.Params.ByName("network-id")

	collection, flawError := entities.GetEntitiesFromPostgres(networkAccountEntity)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworkAccounts")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = nil
	state.Response.Body = collection
	state.Response.Ok()
}
