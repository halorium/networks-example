package identity

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/serializers"

	"golang.org/x/crypto/bcrypt"
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
	Head *Head `json:"head" unmarshal:"head" db:"head"`
	Body *Body `json:"body" unmarshal:"body" db:"body"`
}

type Head struct {
	CreatedAt        string `json:"created-at"        unmarshal:"created-at"        db:"created-at"`
	Etag             string `json:"etag"              unmarshal:"etag"              db:"etag"`
	IdentityID       string `json:"identity-id"       unmarshal:"identity-id"       db:"identity-id"`
	VariationVersion string `json:"variation-version" unmarshal:"variation-version" db:"variation-version"`
}

type Body struct {
	Name       string      `json:"name"       unmarshal:"name"       db:"name"`
	Email      string      `json:"email"      unmarshal:"email"      db:"email"`
	Variant    string      `json:"variant"    unmarshal:"variant"    db:"variant"`
	Variations *Variations `json:"variations" unmarshal:"variations" db:"variations"`
}

type Variations struct {
	Password     string `json:"-" unmarshal:"password"      db:"-"`
	PasswordHash string `json:"-" unmarshal:"password-hash" db:"password-hash"`
	Salt         string `json:"-" unmarshal:"salt"          db:"salt"`
}

func (e *Entity) ValidPassword(password string) bool {
	hashBytes := []byte(e.Body.Variations.PasswordHash)
	passwordBytes := []byte(e.Body.Variations.Salt + password)
	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)

	return err == nil
}

func (e *Entity) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return flaw.New("wrong type").Wrap("cannot Scan")
	}

	err := serializers.JSONUnmarshalTag("db", source, e)

	if err != nil {
		return err
	}

	return nil
}

// Prefix implements Entity interface
func (e *Entity) Prefix() string {
	return "/identities"
}

// ID implements Entity interface
func (e *Entity) ID() string {
	return e.Head.IdentityID
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
	return serializers.JSONMarshalTag("db", e)
}

// BodyJSON implements Entity interface
func (e *Entity) BodyJSON() []byte {
	return serializers.JSONCompactSerializer(e.Body)
}
