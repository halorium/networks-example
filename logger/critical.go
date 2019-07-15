package logger

import (
	"github.com/halorium/networks-example/color"
)

func Critical(tag string, message interface{}) {
	send(
		"critical",
		color.LightRed,
		tag,
		message,
	)
}
