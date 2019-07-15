package main

import "github.com/aws/aws-sdk-go/service/sqs"

type state struct {
	sqsClient       *sqs.SQS
	branchName      string
	queueURL        string
	redriveQueueURL string
	message         *sqs.Message
}
