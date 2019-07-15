package notifier

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/try"
)

func (n *Notifier) Notify(topicName string, message interface{}) *flaw.Error {
	logger.Debug("notifier-notifier-notify", message)

	topicName = envvars.StackName + "-" + topicName

	if n.Topics[topicName] == nil {
		n.addTopic(topicName)
	}

	messageJSON := serializers.JSONCompactSerializer(message)

	snsPublishMessage := &SNSPublishMessage{
		Default: string(messageJSON),
		SQS:     string(messageJSON),
	}

	snsJSON := serializers.JSONCompactSerializer(snsPublishMessage)

	params := &sns.PublishInput{
		TopicArn:         aws.String(n.Topics[topicName].ARN),
		MessageStructure: aws.String("json"),
		Message:          aws.String(string(snsJSON)),
	}

	try := try.New()

	for {
		awsClient := awsclient.FromPool()

		_, err := awsClient.SNS.Publish(params)

		awsclient.ToPool(awsClient)

		flawError := flaw.From(err)

		if flawError != nil {
			if try.Again("notifier-notify-retry", flawError) {
				continue
			}

			return flawError.Wrap("cannot Notify")
		}

		return nil
	}
}
