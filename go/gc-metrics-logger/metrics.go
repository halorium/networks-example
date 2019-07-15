package gcmetricslogger

import (
	"runtime"
	"time"
)

type Metrics struct {
	stopChannel         chan bool
	duration            time.Duration
	memStats            runtime.MemStats
	memStatsTranslation memStatsTranslation
}

func newMetrics(duration time.Duration) *Metrics {
	return &Metrics{
		stopChannel: make(chan bool),
		duration:    duration,
		memStats:    runtime.MemStats{},
	}
}
