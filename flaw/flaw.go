package flaw

import (
	"encoding/json"
	"fmt"
)

type Flaw interface {
	error
	fmt.Stringer
	json.Marshaler
	Wrap(message string) *Error
}
