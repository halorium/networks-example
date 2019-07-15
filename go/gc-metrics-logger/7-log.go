package gcmetricslogger

import "github.com/halorium/networks-example/logger"

func output(metrics *Metrics) {
	logger.Info("gc-metrics", metrics.memStatsTranslation)
}
