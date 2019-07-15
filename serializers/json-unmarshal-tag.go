package serializers

import (
	jsoniter "github.com/halorium/json-iterator"
	"github.com/halorium/networks-example/flaw"
)

func JSONUnmarshalTag(tag string, data []byte, entity interface{}) *flaw.Error {
	err := jsoniter.Config{Tag: tag}.Froze().Unmarshal(data, entity)

	if err != nil {
		return flaw.From(err).Wrap("cannot JSONUnmarshalTag")
	}

	return nil
}
