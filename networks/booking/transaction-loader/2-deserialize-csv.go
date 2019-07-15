package booking

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
)

func deserializeCSV(state *state.State) *flaw.Error {
	logger.Debug("booking-deserialize-csv", state)

	var res []transaction

	flawError := serializers.CSVDeserializer(state.VariantReadCloser, &res)

	if flawError != nil {
		if flawError.Error() == "empty csv file given" {
			logger.Debug("booking-deserialize-csv-empty", state)
			return nil
		}

		return flawError.Wrap("cannot deserializeCSV")
	}

	state.VariantResponse = res

	return iterate(state)
}
