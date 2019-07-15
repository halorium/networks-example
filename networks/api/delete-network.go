package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func deleteNetwork(state *state.State) {
	logger.Debug("delete-network", state)

	state.Network = network.NewEntity()
	state.Network.Head.NetworkID = state.Request.Params.ByName("network-id")

	flawError := entities.DeleteEntityFromPostgres(state.Network)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot deleteNetwork")

		if flawError.Error() == "subordinate entities exist" {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("network has network-accounts")
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
