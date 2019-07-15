package serializers

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/halorium/networks-example/flaw"
)

func CSVDeserializer(readCloser io.ReadCloser, deserializedEntity interface{}) *flaw.Error {
	defer readCloser.Close()

	err := gocsv.Unmarshal(readCloser, deserializedEntity)

	if err != nil {
		return flaw.From(err).Wrap("cannot CSVDeserializer")
	}

	return nil
}
