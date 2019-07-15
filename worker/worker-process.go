package worker

import (
	"github.com/halorium/networks-example/go/gc-metrics-logger"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/profiler"
)

func (w *Worker) Process() {
	profiler.ListenAndServe()

	gcmetricslogger.Start()

	for _, queue := range w.Queues {
		go notifier.Receive(
			notifier.State{
				QueueURL:        queue.URL,
				Handler:         queue.Handler,
				WaitTimeSeconds: 20,
				MaxMessages:     10,
			},
		)
	}

	// keep the process running
	select {}
}
