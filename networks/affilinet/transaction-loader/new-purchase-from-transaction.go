package affilinet

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
	newPurchase.Head.PurchaseID = transaction.TransactionID
	newPurchase.Head.CreatedAt = stamps.New().String()

	newPurchase.Body.ClickID = transaction.SubID
	newPurchase.Body.Status = statusOf(transaction)

	newPurchase.Body.Currency = state.NetworkAccount.Body.Variations.Currency

	timestamp, flawError := stamps.ParseStamp(transaction.RegistrationDate, stamps.YYYYhMMhDDTHHcMMcSS)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.Date = timestamp.String()

	saleAmount, flawError := currency.NewAmountFromString(transaction.NetPrice, newPurchase.Body.Currency)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.SaleAmount = saleAmount.Value.String()

	commissionAmount, flawError := currency.NewAmountFromString(transaction.PublisherCommission, newPurchase.Body.Currency)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.CommissionAmount = commissionAmount.Value.String()

	newPurchase.Body.OrderID = transaction.BasketInfo.BasketID
	newPurchase.Body.Variant = Variant
	newPurchase.Body.Variations = transaction

	return newPurchase, nil
}
