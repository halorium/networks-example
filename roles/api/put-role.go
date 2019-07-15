package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/roles/role"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func putRole(state *state.State) {
	logger.Debug("put-role", state)

	state.Role = role.NewEntity()

	state.Role.Head.RoleID = state.Request.Params.ByName("role-id")
	state.Role.Head.CreatedAt = state.CreatedAt.String()

	flawError := serializers.JSONDeserializer(state.Request.Body, state.Role.Body)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putRole")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := role.Validate(state.Role)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	state.Role.SetEtagHeader()

	// put into entities table
	flawError = entities.PutEntityInPostgres(state.Role)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putRole")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = state.Role.Head
	state.Response.Body = state.Role.Body
	state.Response.Ok()
}
