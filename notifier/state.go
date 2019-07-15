package notifier

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/flaw"
)

type State struct {
	QueueURL        string                     `json:"queue-url"`
	Handler         func(*Message) *flaw.Error `json:"-"`
	WaitTimeSeconds int                        `json:"wait-time-seconds"`
	MaxMessages     int                        `json:"max-messages"`
	Response        *sqs.ReceiveMessageOutput  `json:"-"`
	Message         *Message                   `json:"message"`
	FlawError       *flaw.Error                `json:"flaw-error"`
}
