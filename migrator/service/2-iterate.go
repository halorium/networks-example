package main

import (
	"regexp"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/stamps"
)

func iterate(state *State) {
	logger.Debug("migrator-iterate", state)

	IDRegexp := regexp.MustCompile(`^\d+`)

	for _, state.File = range state.Files {
		state.CreatedAt = stamps.New().String()
		state.FileName = state.File.Name()
		state.FilepathName = envvars.MigrationsPath + "/" + state.FileName
		state.FileID = IDRegexp.FindString(state.FileName)

		filter(state)
	}
}
