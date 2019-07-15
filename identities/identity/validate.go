package identity

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(entity *Entity) *validation.Entity {
	logger.Debug("identity-validate", entity)

	validationEntity := validation.New()

	if entity.Body.Name == "" {
		validationEntity.Add("name is required")
	}

	if entity.Body.Email == "" {
		validationEntity.Add("email is required")
	}

	if entity.Body.Variant == "" {
		validationEntity.Add("variant is required")
	}

	if entity.Body.Variations.Password == "" {
		validationEntity.Add("variations.password is required")
	}

	return validationEntity
}
