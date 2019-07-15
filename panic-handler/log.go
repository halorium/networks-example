package panichandler

import (
	"github.com/halorium/networks-example/logger"
)

func Log(tag string) {
	recovered := recover()

	if recovered != nil {
		logger.Panic(tag, recovered)
	}
}
