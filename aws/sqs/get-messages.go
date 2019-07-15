package awssqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/flaw"
)

func GetMessages(sqsClient *sqs.SQS, queueURL string) ([]*sqs.Message, *flaw.Error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	}

	res, err := sqsClient.ReceiveMessage(params)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot GetMessages")
	}

	return res.Messages, nil
}
