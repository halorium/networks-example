package notifier

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/panic-handler"
	"github.com/halorium/networks-example/sleep"
	"github.com/halorium/networks-example/try"
)

func Receive(state State) {
	defer panichandler.Log("notifier-receive")

	statePointer := &state

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(statePointer.QueueURL),
		WaitTimeSeconds:     aws.Int64(int64(statePointer.WaitTimeSeconds)),
		MaxNumberOfMessages: aws.Int64(int64(statePointer.MaxMessages)),
	}

	// ./crank runs faster with thundering herds :-)
	if envvars.NotMocked {
		sleep.ForAwhile(time.Duration(statePointer.WaitTimeSeconds) * time.Second)
	}

	for {
		// reset these for state each time
		statePointer.FlawError = nil
		statePointer.Message = nil
		statePointer.Response = nil

		logger.Debug("notifier-receive", statePointer)

		try := try.New()

		for {
			awsClient := awsclient.FromPool()

			res, err := awsClient.SQS.ReceiveMessage(params)

			awsclient.ToPool(awsClient)

			flawError := flaw.From(err)

			if flawError != nil {
				if try.Again("notifier-receive-message-retry", flawError) {
					continue
				}

				statePointer.FlawError = flawError

				logger.Critical("notifier-receive", statePointer)

				continue
			}

			statePointer.Response = res

			break
		}

		logger.Info("notifier-received-sqs-messages", statePointer)

		iterate(statePointer)
	}
}
