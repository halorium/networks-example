package logger

import "github.com/halorium/networks-example/color"

func Info(tag string, message interface{}) {
	send("info", color.LightYellow, tag, message)
}
