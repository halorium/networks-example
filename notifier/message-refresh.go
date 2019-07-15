package notifier

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/stamps"
)

func (message *Message) Refresh() *flaw.Error {
	// ./crank cannot ChangeMessageVisibility
	if envvars.Mocked {
		return nil
	}

	if message.RefreshedAt.SecondsSince() < 15 {
		return nil
	}

	params := &sqs.ChangeMessageVisibilityInput{
		QueueUrl:          message.QueueURL,
		ReceiptHandle:     message.ReceiptHandle,
		VisibilityTimeout: aws.Int64(30),
	}

	awsClient := awsclient.FromPool()

	_, err := awsClient.SQS.ChangeMessageVisibility(params)

	awsclient.ToPool(awsClient)

	if err != nil {
		if strings.Contains(err.Error(), "ReceiptHandle is invalid.") {
			return nil
		}

		return flaw.From(err).Wrap("cannot Refresh")
	}

	message.RefreshedAt = stamps.New()

	return nil
}
