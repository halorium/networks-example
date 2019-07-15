package networkaccount

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(entity *Entity) *validation.Entity {
	logger.Debug("network-account-validate", entity)

	validationEntity := validation.New()

	if entity.Body.Name == "" {
		validationEntity.Add("name is required")
	}

	if entity.Body.Variant == "" {
		validationEntity.Add("variant is required")
	}

	if entity.Body.Variations.Currency == "" {
		validationEntity.Add("variations.currency is required")
	}

	return validationEntity
}
