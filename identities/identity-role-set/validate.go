package identityroleset

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(entity *Entity) *validation.Entity {
	logger.Debug("identity-role-set-validate", entity)

	validationEntity := validation.New()

	// TODO we should make sure it's a vailid role in the DB

	return validationEntity
}
