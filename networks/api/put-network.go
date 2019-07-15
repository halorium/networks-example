package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/uuid"
	"github.com/halorium/networks-example/validation"
)

func putNetwork(state *state.State) {
	logger.Debug("put-network", state)

	state.Network = network.NewEntity()

	state.Network.Head.NetworkID = state.Request.Params.ByName("network-id")
	state.Network.Head.CreatedAt = state.CreatedAt.String()

	flawError := serializers.JSONDeserializer(state.Request.Body, state.Network.Body)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putNetwork")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := network.Validate(state.Network)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	state.Network.Head.VariationVersion = uuid.NewHash(
		state.Network.Body.Variations.APIHost,
		state.Network.Body.Variations.DaysPerRequest,
		state.Network.Body.Variations.DaysToLoad,
		state.Network.Body.Variations.RateInterval,
		state.Network.Body.Variations.RateQuantity,
	)

	state.Network.Head.Etag = uuid.NewHash(string(serializers.JSONCompactSerializer(state.Network.Body)))

	flawError = entities.PutEntityInPostgres(state.Network)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putNetwork")

		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = state.Network.Head
	state.Response.Body = state.Network.Body
	state.Response.Ok()
}
