package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/roles/role"
	"github.com/halorium/networks-example/state"
)

func getRole(state *state.State) {
	logger.Debug("get-role", state)

	state.Role = role.NewEntity()

	state.Role.Head.RoleID = state.Request.Params.ByName("role-id")

	flawError := entities.GetEntityFromPostgres(state.Role)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getRole")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.Role.Head
	state.Response.Body = state.Role.Body
	state.Response.Ok()
}
