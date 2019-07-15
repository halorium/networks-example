package notifier

import (
	"encoding/json"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/stamps"
)

func iterate(state *State) {
	for _, sqsMessage := range state.Response.Messages {
		// reset these for state each time
		state.FlawError = nil
		state.Message = nil

		sqsMessageBody := &SQSMessageBody{}

		err := json.Unmarshal([]byte(*sqsMessage.Body), sqsMessageBody)

		if err != nil {
			state.FlawError = flaw.From(err).Wrap("cannot iterate")

			logger.Critical("notifier-iterate-first-unmarshal", state)

			continue
		}

		state.Message = &Message{
			QueueURL:      &state.QueueURL,
			ReceiptHandle: sqsMessage.ReceiptHandle,
			RefreshedAt:   stamps.New(),
		}

		err = json.Unmarshal([]byte(sqsMessageBody.Message), &state.Message.Body)

		if err != nil {
			state.FlawError = flaw.From(err).Wrap("cannot json Unmarshal")

			logger.Critical("notifier-iterate-second-unmarshal", state)

			continue
		}

		dispatch(state)
	}
}
