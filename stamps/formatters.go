package stamps

import (
	"fmt"
	"strconv"
	"time"
)

// h = "-" hyphen
// f = "/" forward slash
// c = ":" colon
// s = " " space
// p = "." period
// T = "T"
// Z = "Z"
// zzzz = "0700"   numerical zone
// zzczz = "07:00" numerical zone with a colon

const DDfMMfYY = "02/01/06"
const DDfMMfYYYY = "02/01/2006"
const DDpMMpYYYY = "02.01.2006"
const DDhMonhYYYY = "02-Jan-2006"
const MMfDDfYY = "01/02/06"
const YYYYhMMhDD = "2006-01-02"
const YYYYhMonhDD = "2006-Jan-02"

const HHMM = "1504"
const HHcMMcSS = "15:04:05"

const DDfMMfYYYYsHHcMMcSS = "02/01/2006 15:04:05"
const YYYYhMMhDDsHHcMMcSS = "2006-01-02 15:04:05"
const YYYYhMMhDDsHHcMMcSSsMST = "2006-01-02 15:04:05 MST"
const YYYYhMMhDDTHHcMMcSS = "2006-01-02T15:04:05"
const YYYYhMMhDDTHHcMMcSSZ = "2006-01-02T15:04:05Z"
const YYYYhMMhDDTHHcMMcSSZzzzz = "2006-01-02T15:04:05Z0700"
const YYYYhMMhDDsHHcMMcSSshzzzz = "2006-01-02 15:04:05 -0700"
const YYYYhMMhDDsHHcMMcSSZhzzzz = "2006-01-02T15:04:05-0700"

const YYYYhMMhDDTHHcMMcSSp000000Z = "2006-01-02T15:04:05.000000Z" // OutputFormat (27 characters)
const StandardFormat = YYYYhMMhDDTHHcMMcSSp000000Z

const YYYYhMMhDDTHHcMMcSShzzczz = "2006-01-02T15:04:05-07:00"
const YYYYhMMhDDTHHcMMcSSp999999hzzczz = "2006-01-02T15:04:05.999999-07:00"
const YYYYhMMhDDTHHcMMcSSZzzczz = time.RFC3339                             // "2006-01-02T15:04:05Z07:00" (RFC3339)
const YYYYhMMhDDTHHcMMcSSp999999Zzzzz = "2006-01-02T15:04:05.999999Z07:00" // "2006-01-02T15:04:05.999999Z07:00" (RFC3339Micro)
const YYYYhMMhDDTHHcMMcSSp999999999Zzzzz = time.RFC3339Nano                // "2006-01-02T15:04:05.999999999Z07:00" (RFC3339Nano)

// Implement the Stringer Interface
func (t *Timestamp) String() string {
	return t.Format(YYYYhMMhDDTHHcMMcSSp000000Z)
}

func (t *Timestamp) YYYY() string {
	return strconv.Itoa(t.Year()) // "2006"
}

func (t *Timestamp) MM() string {
	return fmt.Sprintf("%02d", int(t.Month())) // "02"
}

func (t *Timestamp) DD() string {
	return fmt.Sprintf("%02d", t.Day()) // "02"
}

func (t *Timestamp) LongMonth() string {
	return t.Month().String() // "January"
}

func (t *Timestamp) ShortMonth() string {
	return t.LongMonth()[0:3] // "Jan"
}

func (t *Timestamp) BeginningOfDay() string {
	return t.Format(YYYYhMMhDD) + "T00:00:00.000000Z"
}

func (t *Timestamp) EndOfDay() string {
	return t.Format(YYYYhMMhDD) + "T23:59:59.999999Z"
}
