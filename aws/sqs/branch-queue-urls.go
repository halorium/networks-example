package awssqs

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/flaw"
)

func QueueURLs(sqsClient *sqs.SQS, branchName string) (map[string]string, *flaw.Error) {
	listQueuesInput := &sqs.ListQueuesInput{
		QueueNamePrefix: aws.String(branchName),
	}

	listQueuesOutput, err := sqsClient.ListQueues(listQueuesInput)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot QueueURLs")
	}

	queueURLs := map[string]string{}

	for _, queueURL := range listQueuesOutput.QueueUrls {
		if strings.HasSuffix(*queueURL, "-redrive") {
			continue
		}

		queueURLs[*queueURL] = *queueURL + "-redrive"
	}

	return queueURLs, nil
}
