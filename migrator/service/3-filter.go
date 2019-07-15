package main

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func filter(state *State) {
	logger.Debug("migrator-filter", state)

	count := 0

	err := state.PGClient.QueryRow(
		`SELECT count(*)
			 FROM migrations
			 WHERE id = $1`,
		state.FileID,
	).Scan(&count)

	// the following if statement is very tricky, dragons be thar!
	// If you find yourself thinking "can't we skip the above SQL
	// for FileID 000?"" you're probably not fully considering
	// non-virgin runs... :-)

	if err != nil &&
		state.FileID != "000" &&
		err.Error() != "pq: relation \"migrations\" does not exist" {
		state.FlawError = flaw.From(err).Wrap("cannot filter")

		logger.Panic("migrator-filter-err", state)
	}

	if count == 0 {
		execute(state)
	} else {
		logger.Info("migrator-filter-skip", state)
	}
}
