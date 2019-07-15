package networkaccount

import (
	"fmt"
	"net/url"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/serializers"
)

func NewEntity() *Entity {
	return &Entity{
		Head: &Head{},
		Body: &Body{
			Variations: &Variations{},
		},
	}
}

type Entity struct {
	Head *Head `json:"head"`
	Body *Body `json:"body"`
}

type Head struct {
	CreatedAt        string `json:"created-at"`
	Etag             string `json:"etag"`
	NetworkAccountID string `json:"network-account-id"`
	NetworkID        string `json:"network-id"`
	VariationVersion string `json:"variation-version"`
}

type Body struct {
	Name       string      `json:"name"`
	Variant    string      `json:"variant"`
	Variations *Variations `json:"variations,omitempty"`
}

type Variations struct {
	AdditionalID string `json:"additional-id,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Password     string `json:"password,omitempty"`
	Token        string `json:"token,omitempty"`
	Username     string `json:"username,omitempty"`
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

// Prefix implements Entity interface
func (e *Entity) Prefix() string {
	return fmt.Sprintf("/networks/%s/accounts", url.PathEscape(e.Head.NetworkID))
}

// ID implements Entity interface
func (e *Entity) ID() string {
	return e.Head.NetworkAccountID
}

// Etag implements Entity interface
func (e *Entity) Etag() string {
	return e.Head.Etag
}

// CreatedAt implements Entity interface
func (e *Entity) CreatedAt() string {
	return e.Head.CreatedAt
}

// DBJSON implements Entity interface
func (e *Entity) DBJSON() []byte {
	return serializers.JSONCompactSerializer(e)
}

// BodyJSON implements Entity interface
func (e *Entity) BodyJSON() []byte {
	return serializers.JSONCompactSerializer(e.Body)
}
