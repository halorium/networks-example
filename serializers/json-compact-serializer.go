package serializers

import (
	"encoding/json"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/halt"
)

func JSONCompactSerializer(object interface{}) []byte {
	serialized, err := json.Marshal(object)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return append(serialized, "\n"...)
}
