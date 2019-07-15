package main

import (
	"os/exec"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func execute(state *State) {
	logger.Debug("migrator-execute", state)

	cmd := exec.Command(state.FilepathName)

	loggerEntry, err := cmd.Output()

	if err != nil {
		state.FlawError = flaw.From(err).Wrap(string(loggerEntry)).Wrap("cannot execute")
		logger.Panic("migrator-execute-err", state)
	}

	record(state)
}
