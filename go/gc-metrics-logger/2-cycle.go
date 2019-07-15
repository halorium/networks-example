package gcmetricslogger

import (
	"time"

	"github.com/halorium/networks-example/sleep"
)

func cycle(metrics *Metrics) {
	sleep.ForAwhile(metrics.duration)

	for {
		select {
		case <-metrics.stopChannel:
			return
		case <-time.After(metrics.duration):
			collectMemStats(metrics)
		}
	}
}
