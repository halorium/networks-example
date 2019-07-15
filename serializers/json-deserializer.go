package serializers

import (
	"encoding/json"
	"io"

	"github.com/halorium/networks-example/flaw"
)

func JSONDeserializer(readCloser io.ReadCloser, deserializedEntity interface{}) *flaw.Error {
	defer readCloser.Close()

	err := json.NewDecoder(readCloser).Decode(deserializedEntity)

	if err != nil {
		return flaw.From(err).Wrap("cannot JSONDeserializer")
	}

	return nil
}
