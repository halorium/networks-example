package main

import (
	"io/ioutil"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func list(state *State) {
	logger.Debug("migrator-list", state)

	migrationFiles, err := ioutil.ReadDir(envvars.MigrationsPath)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot list")

		logger.Panic("migrator-list-err", state)
	}

	state.Files = migrationFiles

	iterate(state)
}
