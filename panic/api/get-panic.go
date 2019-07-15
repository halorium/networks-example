package main

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func getPanic(state *state.State) {
	logger.Debug("get-panic", state)

	logger.Panic("get-panic", "GET /panic crashed")
}
