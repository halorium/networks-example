package logger

import "github.com/halorium/networks-example/halt"

func Panic(tag string, message interface{}) {
	Critical(tag, message)

	halt.PanicWith("halted requested via logger.Panic")
}
