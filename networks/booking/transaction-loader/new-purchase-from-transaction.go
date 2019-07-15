package booking

import (
	"github.com/halorium/networks-example/currency"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/stamps"
	"github.com/halorium/networks-example/state"
)

func newPurchaseFromTransaction(state *state.State, transaction transaction) (*networkaccountpurchase.Entity, *flaw.Error) {
	newPurchase := networkaccountpurchase.NewEntity()

	newPurchase.Head.NetworkID = state.Network.Head.NetworkID
	newPurchase.Head.NetworkAccountID = state.NetworkAccount.Head.NetworkAccountID
	newPurchase.Head.CreatedAt = stamps.New().String()

	newPurchase.Head.PurchaseID = transaction.BookNumber

	newPurchase.Body.ClickID = transaction.Label

	var flawError *flaw.Error

	newPurchase.Body.Status, flawError = statusOf(transaction)
	if flawError != nil {
		return nil, flawError
	}

	newPurchase.Body.Currency = state.NetworkAccount.Body.Variations.Currency

	timestamp, flawError := stamps.ParseStamp(transaction.Booked, stamps.YYYYhMMhDDsHHcMMcSS)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.Date = timestamp.String()

	saleAmount, flawError := currency.NewAmountFromString(transaction.Commission, newPurchase.Body.Currency)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.SaleAmount = saleAmount.Value.String()

	commissionAmount, flawError := currency.NewAmountFromString(transaction.Fee, newPurchase.Body.Currency)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.CommissionAmount = commissionAmount.Value.String()

	newPurchase.Body.OrderID = transaction.BookNumber

	newPurchase.Body.Variant = Variant
	newPurchase.Body.Variations = transaction

	return newPurchase, nil
}
