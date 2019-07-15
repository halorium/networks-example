package main

import (
	"io/ioutil"

	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/identities/identity"
	"github.com/halorium/networks-example/identities/identity-role-set"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/postgres"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/uuid"
	"github.com/halorium/networks-example/validation"
)

func postIdentity(state *state.State) {
	logger.Debug("post-identity", state)

	state.Identity = identity.NewEntity()

	state.Identity.Head.IdentityID = uuid.NewRandom()
	state.Identity.Head.CreatedAt = state.CreatedAt.String()

	body, err := ioutil.ReadAll(state.Request.Body)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot postIdentity")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("cannot read body")
		state.Response.BadRequest()
		return
	}

	flawError := serializers.JSONUnmarshalTag("unmarshal", body, state.Identity.Body)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot postIdentity")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := identity.Validate(state.Identity)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	state.Identity.Head.VariationVersion = uuid.NewHash(
		state.Identity.Body.Variations.Password,
		state.Identity.Body.Variations.PasswordHash,
		state.Identity.Body.Variations.Salt,
	)

	state.Identity.Head.Etag = uuid.NewHash(string(serializers.JSONCompactSerializer(state.Identity.Body)))

	pgClient := postgres.ClientFromPool()
	defer postgres.ClientToPool(pgClient)
	pgTxn, err := pgClient.Begin()

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot postIdentity")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	// validate email does not exist
	existingIdentity := identity.NewEntity()
	emailExists := true
	flawError = entities.GetEntityFromPostgresByBodyEmailWithTransaction(existingIdentity, state.Identity.Body.Email, pgTxn)

	if flawError != nil {
		if flawError.Error() == "not found" {
			emailExists = false
		} else {
			err := pgTxn.Rollback()

			if err != nil {
				state.FlawError = flawError.Wrap(err.Error()).Wrap("cannot postIdentity")
			}

			state.Response.Head = nil
			state.Response.Body = nil
			state.Response.ServiceUnavailable()
			return
		}
	}

	if emailExists {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error()).Wrap("cannot postIdentity")
		}

		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("email already exists")
		state.Response.Conflict()
		return
	}

	// put identity and roleset into entities table
	flawError = state.Identity.HashPassword()

	if flawError != nil {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error()).Wrap("cannot postIdentity")
		}

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	flawError = entities.PutEntityInPostgresWithTransaction(state.Identity, pgTxn)

	if flawError != nil {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error()).Wrap("cannot postIdentity")

			state.Response.Head = nil
			state.Response.Body = nil
			state.Response.ServiceUnavailable()
			return
		}

		if flawError.Error() == "duplicate entity" {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("duplicate entity")
			state.Response.Conflict()
			return
		}

		state.FlawError = flawError.Wrap("cannot postIdentity")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.IdentityRoleSet = identityroleset.NewEntity()
	state.IdentityRoleSet.Head.IdentityID = state.Identity.Head.IdentityID
	state.IdentityRoleSet.Head.CreatedAt = state.Identity.Head.CreatedAt

	flawError = entities.PutEntityInPostgresWithTransaction(state.IdentityRoleSet, pgTxn)

	if flawError != nil {
		err := pgTxn.Rollback()

		if err != nil {
			state.FlawError = flawError.Wrap(err.Error()).Wrap("cannot postIdentity")

			state.Response.Head = nil
			state.Response.Body = nil
			state.Response.ServiceUnavailable()
			return
		}

		if flawError.Error() == "duplicate entity" {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("duplicate entity")
			state.Response.Conflict()
			return
		}

		state.FlawError = flawError.Wrap("cannot postIdentity")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	err = pgTxn.Commit()

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot postIdentity")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = state.Identity.Head
	state.Response.Body = state.Identity.Body
	state.Response.Created()
}
