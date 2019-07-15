package stamps

import (
	"time"

	"github.com/halorium/networks-example/flaw"
)

func ParseStamp(stamp string, format string) (*Timestamp, *flaw.Error) {
	parsed, err := time.ParseInLocation(format, stamp, time.UTC)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot ParseStamp")
	}

	return &Timestamp{&parsed}, nil
}
