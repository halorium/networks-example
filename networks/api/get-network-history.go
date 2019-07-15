package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network"
	"github.com/halorium/networks-example/state"
)

func getNetworkHistory(state *state.State) {
	logger.Debug("get-network-history", state)

	state.Network = network.NewEntity()

	state.Network.Head.NetworkID = state.Request.Params.ByName("network-id")

	collection, flawError := entities.GetEntityHistoryFromPostgres(state.Network)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworkHistory")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.Network.Head
	state.Response.Body = collection
	state.Response.Ok()
}
