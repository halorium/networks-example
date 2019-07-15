package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/identities/identity"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func getIdentityHistory(state *state.State) {
	logger.Debug("get-identity-history", state)

	state.Identity = identity.NewEntity()

	state.Identity.Head.IdentityID = state.Request.Params.ByName("identity-id")

	collection, flawError := entities.GetEntityHistoryFromPostgres(state.Identity)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getIdentityHistory")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.Identity.Head
	state.Response.Body = collection
	state.Response.Ok()
}
