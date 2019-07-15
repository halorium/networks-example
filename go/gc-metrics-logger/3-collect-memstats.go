package gcmetricslogger

import "runtime"

func collectMemStats(metrics *Metrics) {
	runtime.ReadMemStats(&metrics.memStats)

	translateMemstats(metrics)
}
