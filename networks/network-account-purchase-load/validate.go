package networkaccountpurchaseload

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(entity *Entity) *validation.Entity {
	logger.Debug("network-account-purchase-load-validate", entity)

	validationEntity := validation.New()

	if entity.StartTime == nil {
		validationEntity.Add("start-time must be formatted as 2006-01-02T15:04:05.000000Z")
	}

	if entity.StopTime == nil {
		validationEntity.Add("stop-time must be formatted as 2006-01-02T15:04:05.000000Z")
	}

	return validationEntity
}
