package main

import (
	"github.com/halorium/networks-example/authentication/authorization"
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/identities/identity"
	"github.com/halorium/networks-example/json-web-token"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func postAuthentication(state *state.State) {
	logger.Debug("post-authentication", state)

	email, password, ok := state.Request.BasicAuthCredentials()

	if !ok {
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.BadRequest()
		return
	}

	state.Identity = identity.NewEntity()

	state.Identity.Body.Email = email

	pgClient := postgres.ClientFromPool()
	defer postgres.ClientToPool(pgClient)
	pgTxn, err := pgClient.Begin()

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot postAuthentication")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	flawError := entities.GetEntityFromPostgresByBodyEmailWithTransaction(state.Identity, email, pgTxn)

	if flawError != nil {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error())
		}

		state.FlawError = flawError.Wrap("cannot postAuthentication")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("invalid authentication")
		state.Response.Unauthorized()
		return
	}

	if !state.Identity.ValidPassword(password) {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error())
		}

		state.FlawError = flaw.New("invalid authentication").Wrap("cannot postAuthentication")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("invalid authentication")
		state.Response.Unauthorized()
		return
	}

	state.Authorization = authorization.NewEntity()
	state.Authorization.Head.CreatedAt = state.CreatedAt.String()

	claims := jsonwebtoken.NewClaims()
	claims.DelegatorIdentity = jsonwebtoken.Identity{
		IdentityID: state.Identity.Head.IdentityID,
	}
	claims.DelegatedIdentity = jsonwebtoken.Identity{
		IdentityID: state.Identity.Head.IdentityID,
	}

	JSONWebToken, flawError := jsonwebtoken.New(claims)

	if flawError != nil {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error())
		}

		state.FlawError = flawError.Wrap("cannot postAuthentication")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.BadRequest()
		return
	}

	err = pgTxn.Commit()

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot postAuthentication")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Authorization.Body.Authorization = "Bearer " + JSONWebToken.SignedToken

	state.Response.Head = state.Authorization.Head
	state.Response.Body = state.Authorization.Body
	state.Response.Ok()
}
