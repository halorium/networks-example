package network

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(entity *Entity) *validation.Entity {
	logger.Debug("network-validate", entity)

	validationEntity := validation.New()

	if entity.Body.Name == "" {
		validationEntity.Add("name is required")
	}

	if entity.Body.Variant == "" {
		validationEntity.Add("variant is required")
	}

	if entity.Body.Variations.APIHost == "" {
		validationEntity.Add("variations.api-host is required")
	}

	if entity.Body.Variations.DaysPerRequest == "" {
		validationEntity.Add("variations.days-per-request is required")
	}

	if entity.Body.Variations.DaysToLoad == "" {
		validationEntity.Add("variations.days-to-load is required")
	}

	if entity.Body.Variations.RateInterval == "" {
		validationEntity.Add("variations.rate-interval is required")
	}

	if entity.Body.Variations.RateQuantity == "" {
		validationEntity.Add("variations.rate-quantity is required")
	}

	return validationEntity
}
