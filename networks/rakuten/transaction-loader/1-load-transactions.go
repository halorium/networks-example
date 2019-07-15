package rakuten

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func LoadTransactions(state *state.State) *flaw.Error {
	logger.Debug("rakuten-load-transactions", state)

	return buildRegexp(state)
}
