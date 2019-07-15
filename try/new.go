package try

import "time"

func New() *Try {
	return &Try{
		Start: time.Now(),
		Max:   3,
	}
}
