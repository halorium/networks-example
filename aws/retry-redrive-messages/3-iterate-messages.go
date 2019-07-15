package main

import (
	"time"

	"github.com/halorium/networks-example/aws/sqs"
	"github.com/halorium/networks-example/flaw"
)

func iterateMessages(state *state) *flaw.Error {
	emptyGets := 0

	for {
		messages, flawError := awssqs.GetMessages(state.sqsClient, state.redriveQueueURL)

		if flawError != nil {
			return flawError.Wrap("cannot iterateMessages")
		}

		if len(messages) == 0 {
			emptyGets++

			if emptyGets > 10 {
				break
			}

			time.Sleep(time.Second)

			continue
		}

		for _, message := range messages {
			state.message = message

			flawError = moveMessage(state)

			if flawError != nil {
				return flawError.Wrap("cannot iterateMessages")
			}
		}
	}

	return nil
}
