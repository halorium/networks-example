package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/uuid"
	"github.com/halorium/networks-example/validation"
)

func putNetworkAccount(state *state.State) {
	logger.Debug("put-network-account", state)

	// validate network
	getNetwork(state)

	if state.FlawError != nil {
		state.FlawError = state.FlawError.Wrap("cannot putNetworkAccount")

		if state.FlawError.Error() == "not found" {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("invalid network-id")
			state.Response.BadRequest()
			return
		}

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.NetworkAccount = networkaccount.NewEntity()

	state.NetworkAccount.Head.NetworkID = state.Request.Params.ByName("network-id")
	state.NetworkAccount.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")
	state.NetworkAccount.Head.CreatedAt = state.CreatedAt.String()

	flawError := serializers.JSONDeserializer(state.Request.Body, state.NetworkAccount.Body)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putNetworkAccount")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := networkaccount.Validate(state.NetworkAccount)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	state.NetworkAccount.Head.VariationVersion = uuid.NewHash(
		state.NetworkAccount.Body.Variations.AdditionalID,
		state.NetworkAccount.Body.Variations.Currency,
		state.NetworkAccount.Body.Variations.Password,
		state.NetworkAccount.Body.Variations.Token,
		state.NetworkAccount.Body.Variations.Username,
	)

	state.NetworkAccount.Head.Etag = uuid.NewHash(string(serializers.JSONCompactSerializer(state.NetworkAccount.Body)))

	flawError = entities.PutEntityInPostgres(state.NetworkAccount)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putNetworkAccount")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = state.NetworkAccount.Head
	state.Response.Body = state.NetworkAccount.Body
	state.Response.Ok()
}
