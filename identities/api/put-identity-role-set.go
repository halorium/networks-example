package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/identities/identity-role-set"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func putIdentityRoleSet(state *state.State) {
	logger.Debug("put-identity-role-set", state)

	getIdentity(state)

	if state.FlawError != nil {
		return
	}

	state.IdentityRoleSet = identityroleset.NewEntity()
	state.IdentityRoleSet.Head.IdentityID = state.Request.Params.ByName("identity-id")
	state.IdentityRoleSet.Head.CreatedAt = state.CreatedAt.String()

	flawError := serializers.JSONDeserializer(state.Request.Body, state.IdentityRoleSet.Body)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putIdentityRoleSet")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := identityroleset.Validate(state.IdentityRoleSet)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	state.IdentityRoleSet.SetEtagHeader()

	flawError = entities.PutEntityInPostgres(state.IdentityRoleSet)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putIdentityRoleSet")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = state.IdentityRoleSet.Head
	state.Response.Body = state.IdentityRoleSet.Body
	state.Response.Ok()
}
