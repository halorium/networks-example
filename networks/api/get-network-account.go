package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account"
	"github.com/halorium/networks-example/state"
)

func getNetworkAccount(state *state.State) {
	logger.Debug("get-network-account", state)

	state.NetworkAccount = networkaccount.NewEntity()

	state.NetworkAccount.Head.NetworkID = state.Request.Params.ByName("network-id")
	state.NetworkAccount.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")

	flawError := entities.GetEntityFromPostgres(state.NetworkAccount)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworkAccount")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.NetworkAccount.Head
	state.Response.Body = state.NetworkAccount.Body
	state.Response.Ok()
}
