package main

import (
	"fmt"
	"os"

	"github.com/halorium/networks-example/aws/sqs"
	"github.com/halorium/networks-example/flaw"
)

func iterateQueues(state *state) *flaw.Error {
	queueURLs, flawError := awssqs.QueueURLs(state.sqsClient, state.branchName)

	if flawError != nil {
		return flawError.Wrap("cannot iterateQueues")
	}

	for queueURL, redriveQueueURL := range queueURLs {
		_, err := fmt.Fprintln(os.Stderr, queueURL)

		if err != nil {
			return flaw.From(err).Wrap("cannot iterateQueues")
		}

		state.queueURL = queueURL
		state.redriveQueueURL = redriveQueueURL

		flawError = iterateMessages(state)

		if flawError != nil {
			return flawError.Wrap("cannot iterateQueues")
		}
	}

	return nil
}
