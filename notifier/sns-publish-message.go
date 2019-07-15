package notifier

type SNSPublishMessage struct {
	Default string `json:"default"`
	SQS     string `json:"sqs"`
}
