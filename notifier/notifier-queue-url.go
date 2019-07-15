package notifier

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func (n *Notifier) QueueURL(queueName string) string {
	awsClient := awsclient.FromPool()

	res, err := awsClient.SQS.GetQueueUrl(
		&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)},
	)

	awsclient.ToPool(awsClient)

	if err != nil {
		logger.Panic("notifer-notifier-queue-url", flaw.From(err))
	}

	return *res.QueueUrl
}
