package notifier

import "github.com/halorium/networks-example/stamps"

type Message struct {
	Body          map[string]interface{} `json:"body"`
	QueueURL      *string                `json:"queue-url"`
	ReceiptHandle *string                `json:"receipt-handle"`
	RefreshedAt   *stamps.Timestamp      `json:"refreshed-at"`
}
