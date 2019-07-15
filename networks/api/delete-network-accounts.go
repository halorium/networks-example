package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func deleteNetworkAccounts(state *state.State) {
	logger.Debug("delete-network-accounts", state)

	state.NetworkAccount = networkaccount.NewEntity()
	state.NetworkAccount.Head.NetworkID = state.Request.Params.ByName("network-id")

	flawError := entities.DeleteEntitiesFromPostgres(state.NetworkAccount)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot deleteNetworkAccounts")

		if flawError.Error() == "subordinate entities exist" {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("network-account has network-account-purchases")
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
