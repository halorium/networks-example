package gcmetricslogger

import "time"

type memStatsTranslation struct {
	BuckHashSys         uint64        `json:"bucket-hash-sys"`
	Frees               uint64        `json:"frees"`
	GCCPUFraction       float64       `json:"gc-cpu-fraction"`
	GCSys               uint64        `json:"gc-sys"`
	HeapAlloc           uint64        `json:"heap-alloc"`
	HeapIdle            uint64        `json:"heap-idle"`
	HeapInuse           uint64        `json:"heap-in-use"`
	HeapObjects         uint64        `json:"heap-objects"`
	HeapReleased        uint64        `json:"heap-released"`
	HeapSys             uint64        `json:"heap-sys"`
	LastGC              uint64        `json:"last-gc"`
	LastPauseDuration   time.Duration `json:"last-pause-duration"`
	LastPauseFinishedAt time.Time     `json:"last-pause-finished-at"`
	Lookups             uint64        `json:"lookups"`
	MCacheInuse         uint64        `json:"mcache-in-use"`
	MCacheSys           uint64        `json:"mcache-sys"`
	MSpanInuse          uint64        `json:"mspan-in-use"`
	MSpanSys            uint64        `json:"mspan-sys"`
	Mallocs             uint64        `json:"mallocs"`
	NextGC              uint64        `json:"next-gc"`
	NumForcedGC         uint32        `json:"num-forced-gc"`
	NumGC               uint32        `json:"num-gc"`
	OtherSys            uint64        `json:"other-sys"`
	PauseTotalNs        uint64        `json:"pause-total-ns"`
	StackInuse          uint64        `json:"stack-in-use"`
	StackSys            uint64        `json:"stack-sys"`
	Sys                 uint64        `json:"sys"`
	TotalAlloc          uint64        `json:"total-alloc"`
}
