package stamps

import "time"

func (t *Timestamp) SecondsSince() float64 {
	return time.Since(*t.Time).Seconds()
}
