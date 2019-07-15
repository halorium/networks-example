package gcmetricslogger

import "time"

func setLastPauseDuration(metrics *Metrics) {
	metrics.memStatsTranslation.LastPauseDuration =
		time.Duration(
			metrics.memStats.PauseNs[(metrics.memStats.NumGC+255)%256],
		)

	setLastPauseFinishedAt(metrics)
}
