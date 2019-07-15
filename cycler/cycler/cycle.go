package cycler

import (
	"time"

	"github.com/halorium/networks-example/stamps"
)

type Cycle struct {
	DaysToLoad     int               `json:"days-to-load,omitempty"`
	DaysPerRequest int               `json:"days-per-request,omitempty"`
	StartTime      *stamps.Timestamp `json:"start-time,omitempty"`
	StopTime       *stamps.Timestamp `json:"stop-time,omitempty"`
	Period         string            `json:"period,omitempty"`
	Action         func(*Cycle)      `json:"-"`
	Tag            string            `json:"tag,omitempty"`
	URI            string            `json:"uri,omitempty"`
	Sleep          time.Duration     `json:"sleep,omitempty"`
	RateInterval   string            `json:"rate-interval,omitempty"`
	RateQuantity   string            `json:"rate-quantity,omitempty"`
}
