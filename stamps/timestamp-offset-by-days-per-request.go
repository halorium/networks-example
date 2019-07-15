package stamps

import "time"

func (t *Timestamp) OffsetByDaysPerRequest(daysToLoad int) *Timestamp {
	var ts time.Time

	if daysToLoad > 1 {
		ts = t.AddDate(0, 0, -daysToLoad)
	} else {
		ts = t.AddDate(0, 0, -0)
	}
	return &Timestamp{&ts}
}
