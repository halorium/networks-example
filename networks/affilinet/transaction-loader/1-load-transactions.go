package affilinet

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func LoadTransactions(state *state.State) *flaw.Error {
	logger.Debug("affilinet-load-transactions", state)

	return authenticate(state)
}
