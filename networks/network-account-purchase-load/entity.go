package networkaccountpurchaseload

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/stamps"
)

func NewEntity() *Entity {
	return &Entity{}
}

type Entity struct {
	StartTime *stamps.Timestamp `json:"start-time"`
	StopTime  *stamps.Timestamp `json:"stop-time"`
}

func (e *Entity) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return flaw.New("wrong type").Wrap("cannot Scan")
	}

	flawError := serializers.JSONUnmarshalTag("json", source, e)
	if flawError != nil {
		return flawError.Wrap("cannot Scan")
	}

	return nil
}
