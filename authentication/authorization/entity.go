package authorization

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/serializers"
)

func NewEntity() *Entity {
	return &Entity{
		Head: &Head{},
		Body: &Body{},
	}
}

type Entity struct {
	Head *Head `json:"head"`
	Body *Body `json:"body"`
}

type Head struct {
	CreatedAt string `json:"created-at"`
}

type Body struct {
	Authorization string `json:"authorization"`
}

func (e *Entity) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return flaw.New("wrong type").Wrap("cannot Scan")
	}

	err := serializers.JSONUnmarshalTag("json", source, e)
	if err != nil {
		return flaw.From(err).Wrap("cannot Scan")
	}

	return nil
}
