package gcmetricslogger

func translateMemstats(metrics *Metrics) {
	metrics.memStatsTranslation = memStatsTranslation{
		BuckHashSys:   metrics.memStats.BuckHashSys,
		Frees:         metrics.memStats.Frees,
		GCCPUFraction: metrics.memStats.GCCPUFraction,
		GCSys:         metrics.memStats.GCSys,
		HeapAlloc:     metrics.memStats.HeapAlloc,
		HeapIdle:      metrics.memStats.HeapIdle,
		HeapInuse:     metrics.memStats.HeapInuse,
		HeapObjects:   metrics.memStats.HeapObjects,
		HeapReleased:  metrics.memStats.HeapReleased,
		HeapSys:       metrics.memStats.HeapSys,
		Lookups:       metrics.memStats.Lookups,
		MCacheInuse:   metrics.memStats.MCacheInuse,
		MCacheSys:     metrics.memStats.MCacheSys,
		MSpanInuse:    metrics.memStats.MSpanInuse,
		MSpanSys:      metrics.memStats.MSpanSys,
		Mallocs:       metrics.memStats.Mallocs,
		NextGC:        metrics.memStats.NextGC,
		NumForcedGC:   metrics.memStats.NumForcedGC,
		NumGC:         metrics.memStats.NumGC,
		OtherSys:      metrics.memStats.OtherSys,
		PauseTotalNs:  metrics.memStats.PauseTotalNs,
		StackInuse:    metrics.memStats.StackInuse,
		StackSys:      metrics.memStats.StackSys,
		Sys:           metrics.memStats.Sys,
		TotalAlloc:    metrics.memStats.TotalAlloc,
	}

	setLastPauseDuration(metrics)
}
