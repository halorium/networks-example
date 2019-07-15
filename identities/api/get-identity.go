package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/identities/identity"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func getIdentity(state *state.State) {
	logger.Debug("get-identity", state)

	state.Identity = identity.NewEntity()

	state.Identity.Head.IdentityID = state.Request.Params.ByName("identity-id")

	flawError := entities.GetEntityFromPostgres(state.Identity)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getIdentity")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.Identity.Head
	state.Response.Body = state.Identity.Body
	state.Response.Ok()
}
