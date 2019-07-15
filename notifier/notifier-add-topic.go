package notifier

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/halorium/networks-example/aws/client"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func (n *Notifier) addTopic(topicName string) {
	topic := &Topic{
		Name: topicName,
	}

	awsClient := awsclient.FromPool()

	res, err := awsClient.SNS.ListTopics(&sns.ListTopicsInput{})

	awsclient.ToPool(awsClient)

	if err != nil {
		logger.Panic("notifier-notifier-add-topic", flaw.From(err))
	}

	for _, snsTopic := range res.Topics {
		if strings.HasSuffix(*snsTopic.TopicArn, topicName) {
			topic.ARN = *snsTopic.TopicArn
			break
		}
	}

	n.Topics[topicName] = topic
}
