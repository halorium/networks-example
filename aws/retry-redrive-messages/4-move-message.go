package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/flaw"
)

func moveMessage(state *state) *flaw.Error {
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: state.message.Body,
		QueueUrl:    aws.String(state.queueURL),
	}

	_, err := state.sqsClient.SendMessage(sendMessageInput)

	if err != nil {
		return flaw.From(err).Wrap("cannot moveMessage")
	}

	deleteMessageInput := &sqs.DeleteMessageInput{
		QueueUrl:      &state.redriveQueueURL,
		ReceiptHandle: state.message.ReceiptHandle,
	}

	_, err = state.sqsClient.DeleteMessage(deleteMessageInput)

	if err != nil {
		return flaw.From(err).Wrap("cannot moveMessage")
	}

	// TODO use logger in the future, requires tons of env vars to be set
	fmt.Println(*state.message.Body)

	return nil
}
