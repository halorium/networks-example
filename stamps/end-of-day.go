package stamps

import (
	"github.com/halorium/networks-example/logger"
)

func EndOfDay(stamp *Timestamp) *Timestamp {
	stampString := stamp.Format(YYYYhMMhDD) + "T23:59:59.999999Z"
	ts, flawError := ParseStamp(stampString, YYYYhMMhDDTHHcMMcSSp999999Zzzzz)
	if flawError != nil {
		logger.Panic("stamps-end-of-day", flawError)
	}
	return ts
}
