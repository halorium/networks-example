package awshttpsigningclient

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func authentication(state *State) {
	if !fillFromEnvKeys(state) &&
		!fillFromTaskRole(state) {
		logger.Panic(
			"awshttpsigningclient-authentication",
			flaw.New("no credentials available").Wrap("cannot authenticate"),
		)
	}

	requestTimes(state)
}
