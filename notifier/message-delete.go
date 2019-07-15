package notifier

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/try"
)

func (message *Message) Delete() {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      message.QueueURL,
		ReceiptHandle: message.ReceiptHandle,
	}

	try := try.New()

	for {
		awsClient := awsclient.FromPool()

		_, err := awsClient.SQS.DeleteMessage(params)

		awsclient.ToPool(awsClient)

		if err != nil {
			if try.Again("notifier-message-delete-retry", flaw.From(err)) {
				continue
			}
		}

		break
	}
}
