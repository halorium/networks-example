package logger

import "github.com/halorium/networks-example/color"

func Warn(tag string, message interface{}) {
	send("warn", color.LightMagenta, tag, message)
}
