package serializers

import (
	"bytes"
	"encoding/json"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/halt"
)

func JSONPrettySerializer(object interface{}) []byte {
	buffer := bytes.NewBuffer(nil)

	encoder := json.NewEncoder(buffer)

	encoder.SetIndent("", "  ")

	err := encoder.Encode(object)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return buffer.Bytes()
}
