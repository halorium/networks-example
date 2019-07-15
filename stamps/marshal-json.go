package stamps

import (
	"fmt"
	"time"
)

// MarshalJSON implements the json.Marshaler interface.
func (stamp *Timestamp) MarshalJSON() ([]byte, error) {
	stampString := fmt.Sprintf("\"%s\"", stamp.String())
	return []byte(stampString), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (stamp *Timestamp) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	timeTime, err := time.Parse(`"`+YYYYhMMhDDTHHcMMcSSp000000Z+`"`, string(data))
	stamp.Time = &timeTime
	return err
}
