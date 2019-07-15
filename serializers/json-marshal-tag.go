package serializers

import (
	jsoniter "github.com/halorium/json-iterator"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/halt"
)

func JSONMarshalTag(tag string, entity interface{}) []byte {
	serialized, err := jsoniter.Config{Tag: tag}.Froze().Marshal(entity)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return append(serialized, "\n"...)
}
