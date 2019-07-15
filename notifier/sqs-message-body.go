package notifier

type SQSMessageBody struct {
	Type             string
	MessageID        string
	TopicArn         string
	Message          string
	Timestamp        string
	SignatureVersion string
	Signature        string
	SigningCertURL   string
	UnsubscribeURL   string
}
