package affiliatewindow

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
)

func deserializeJSON(state *state.State) *flaw.Error {
	logger.Debug("affiliate-window-deserialize-json", state)

	var res []rawTransaction

	flawError := serializers.JSONDeserializer(state.VariantReadCloser, &res)

	if flawError != nil {
		return flawError.Wrap("cannot deserializeJSON")
	}

	state.VariantResponse = res

	return iterate(state)
}
