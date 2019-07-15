package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network"
	"github.com/halorium/networks-example/state"
)

func getNetworks(state *state.State) {
	logger.Debug("get-networks", state)

	newNetwork := network.NewEntity()

	collection, flawError := entities.GetEntitiesFromPostgres(newNetwork)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworks")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = nil
	state.Response.Body = collection
	state.Response.Ok()
}
