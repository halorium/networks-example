package stamps

import (
	"github.com/halorium/networks-example/logger"
)

func BeginningOfDay(stamp *Timestamp) *Timestamp {
	stampString := stamp.Format(YYYYhMMhDD) + "T00:00:00.000000Z"
	ts, flawError := ParseStamp(stampString, YYYYhMMhDDTHHcMMcSSp999999Zzzzz)
	if flawError != nil {
		logger.Panic("stamps-beginning-of-day", flawError)
	}
	return ts
}
