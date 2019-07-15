package affilinet

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
)

func deserializeXML(state *state.State) *flaw.Error {
	logger.Debug("affilinet-deserialize-xml", state)

	var res response

	flawError := serializers.XMLDeserializer(state.VariantReadCloser, &res)

	if flawError != nil {
		return flawError.Wrap("cannot deserializeXML")
	}

	state.VariantResponse = res

	return iterate(state)
}
