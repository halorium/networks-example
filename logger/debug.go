package logger

import (
	"github.com/halorium/networks-example/color"
)

func Debug(tag string, message interface{}) {
	if InLogDebugMessages != "true" {
		return
	}

	send("debug", color.LightBlack, tag, message)
}
