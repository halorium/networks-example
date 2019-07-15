package serializers

import (
	"encoding/xml"
	"io"

	"github.com/halorium/networks-example/flaw"
)

func XMLDeserializer(readCloser io.ReadCloser, deserializedEntity interface{}) *flaw.Error {
	defer readCloser.Close()

	err := xml.NewDecoder(readCloser).Decode(deserializedEntity)

	if err != nil {
		return flaw.From(err).Wrap("cannot XMLDeserializer")
	}

	return nil
}
