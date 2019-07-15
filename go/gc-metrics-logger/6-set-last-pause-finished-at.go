package gcmetricslogger

import "time"

func setLastPauseFinishedAt(metrics *Metrics) {
	metrics.memStatsTranslation.LastPauseFinishedAt =
		time.Unix(
			0,
			int64(metrics.memStats.PauseEnd[(metrics.memStats.NumGC+255)%256]),
		)

	output(metrics)
}
