package role

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(entity *Entity) *validation.Entity {
	logger.Debug("role-validate", entity)

	validationEntity := validation.New()

	if entity.Body.Name == "" {
		validationEntity.Add("name is required")
	}

	if entity.Body.Description == "" {
		validationEntity.Add("description is required")
	}

	return validationEntity
}
