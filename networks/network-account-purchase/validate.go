package networkaccountpurchase

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/validation"
)

func Validate(networkAccountPurchaseEntity *Entity) *validation.Entity {
	logger.Debug("network-account-purchase-validate", networkAccountPurchaseEntity)

	validationEntity := validation.New()

	if networkAccountPurchaseEntity.Body.Status == "" {
		validationEntity.Add("status is required")
	}

	if networkAccountPurchaseEntity.Body.Currency == "" {
		validationEntity.Add("currency is required")
	}

	if networkAccountPurchaseEntity.Body.Date == "" {
		validationEntity.Add("date is required")
	}

	if networkAccountPurchaseEntity.Body.SaleAmount == "" {
		validationEntity.Add("sale-amount is required")
	}

	if networkAccountPurchaseEntity.Body.CommissionAmount == "" {
		validationEntity.Add("commission-amount is required")
	}

	if networkAccountPurchaseEntity.Body.OrderID == "" {
		validationEntity.Add("order-id is required")
	}

	if networkAccountPurchaseEntity.Body.Variant == "" {
		validationEntity.Add("variant is required")
	}

	if networkAccountPurchaseEntity.Body.Variations == nil {
		validationEntity.Add("variations is required")
	}

	return validationEntity
}
