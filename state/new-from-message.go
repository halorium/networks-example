package state

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/serializers"
)

func NewFromMessage(message *notifier.Message) (*State, *flaw.Error) {
	state := New()

	body := serializers.JSONMarshalTag("json", message.Body)

	flawError := serializers.JSONUnmarshalTag("json", body, state)

	if flawError != nil {
		return nil, flawError
	}

	state.Message = message

	return state, nil
}
