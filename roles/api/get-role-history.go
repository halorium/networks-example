package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/roles/role"
	"github.com/halorium/networks-example/state"
)

func getRoleHistory(state *state.State) {
	logger.Debug("get-role-history", state)

	state.Role = role.NewEntity()

	state.Role.Head.RoleID = state.Request.Params.ByName("role-id")

	collection, flawError := entities.GetEntityHistoryFromPostgres(state.Role)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getRoleHistory")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = state.Role.Head
	state.Response.Body = collection
	state.Response.Ok()
}
