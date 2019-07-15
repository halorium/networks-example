package state

import (
	"github.com/halorium/networks-example/responses"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/stamps"
	"github.com/halorium/networks-example/uuid"
)

func New() *State {
	state := &State{

		// request
		ID:        uuid.NewRandom(),
		CreatedAt: stamps.New(),
	}

	state.Response = responses.NewResponse()
	state.Response.Serializer = serializers.JSONPrettySerializer
	state.Response.State = state // TODO: eliminate cyclical reference!

	return state
}
