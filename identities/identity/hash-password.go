package identity

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (entity *Entity) HashPassword() *flaw.Error {
	if entity.Body.Variations.Password == "" {
		return flaw.New("password must not be empty").Wrap("cannot HashPassword")
	}

	if entity.Body.Variations.Salt != "" && entity.Body.Variations.PasswordHash != "" {
		return flaw.New("password already hashed").Wrap("Cannot HashPassword")
	}

	entity.Body.Variations.Salt = uuid.NewRandom()

	passwordBytes := []byte(entity.Body.Variations.Salt + entity.Body.Variations.Password)
	hashBytes, err := bcrypt.GenerateFromPassword(passwordBytes, 10)

	if err != nil {
		return flaw.From(err).Wrap("cannot HashPassword")
	}

	entity.Body.Variations.PasswordHash = string(hashBytes)

	return nil
}
