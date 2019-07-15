package gcmetricslogger

import (
	"time"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func Start() *Metrics {
	duration, err := time.ParseDuration("150s")

	if err != nil {
		logger.Panic("gc-metrics-logger", flaw.From(err))
	}

	metrics := newMetrics(duration)

	go cycle(metrics)

	return metrics
}
