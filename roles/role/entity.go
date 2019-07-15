package role

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/uuid"
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
	Etag      string `json:"etag"`
	RoleID    string `json:"role-id"`
}

type Body struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Actions     []string `json:"actions"`
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
	return "/roles"
}

// ID implements Entity interface
func (e *Entity) ID() string {
	return e.Head.RoleID
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

// SetEtagHeader sets the header before being saved in the db
// the if block is to handle going into and coming out of the DB
func (e *Entity) SetEtagHeader() {
	if e.Head.Etag == "" {
		e.Head.Etag = uuid.NewHash(string(serializers.JSONCompactSerializer(e.Body)))
	}
}
