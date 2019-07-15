package affiliatewindow

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/uuid"
)

func SetVariationVersion(networkAccountPurchase *networkaccountpurchase.Entity) *flaw.Error {
	body := serializers.JSONMarshalTag("json", networkAccountPurchase.Body.Variations)

	trans := &transaction{}

	flawError := serializers.JSONUnmarshalTag("json", body, trans)

	if flawError != nil {
		return flawError.Wrap("cannot SetVariationVersion")
	}

	networkAccountPurchase.Body.OrderID = trans.OrderRef

	if networkAccountPurchase.Body.OrderID == "" {
		networkAccountPurchase.Body.OrderID = networkAccountPurchase.ID()
	}

	networkAccountPurchase.Head.VariationVersion = uuid.NewHash(
		trans.ID,                            // purchase-id
		networkAccountPurchase.Body.ClickID, // click-id
		trans.AwinCommissionAmount.Amount,   // commission
		trans.AwinSaleAmount.Currency,       // currency
		trans.TransactionDate,               // date
		trans.OrderRef,                      // order-id
		trans.AwinSaleAmount.Amount,         // sale
		trans.CommissionStatus,              // status
	)

	return nil
}
