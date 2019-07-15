package stamps

import "github.com/halorium/networks-example/flaw"

// ParseDateAndTime Convert ("2006-Jan-02", YYYYhMonhDD, "1504", HHMM) to UTC Timestamp
func ParseDateAndTime(dateString, dateFormat, timeString, timeFormat string) (*Timestamp, *flaw.Error) {
	format := dateFormat + "T" + timeFormat
	dateAndTimeString := dateString + "T" + timeString

	return ParseStamp(dateAndTimeString, format)
}
