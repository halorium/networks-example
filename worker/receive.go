package worker

import (
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/notifier"
)

func (w *Worker) Receive(queueName string, handler func(*notifier.Message) *flaw.Error) {
	queueName = envvars.StackName + "-" + queueName

	n := notifier.FromPool()

	url := n.QueueURL(queueName)

	notifier.ToPool(n)

	queue := Queue{
		Name:    queueName,
		Handler: handler,
		URL:     url,
	}

	w.Queues[queueName] = queue
}
