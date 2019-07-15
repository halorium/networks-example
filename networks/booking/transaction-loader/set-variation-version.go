package booking

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

	networkAccountPurchase.Head.VariationVersion = uuid.NewHash(
		trans.BookNumber, // purchase-id
		trans.Label,      // click-id
		trans.Fee,        // commission
		trans.Booked,     // date
		trans.BookNumber, // order-id
		trans.Commission, // sale
	)

	return nil
}
