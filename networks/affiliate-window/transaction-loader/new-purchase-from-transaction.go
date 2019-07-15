package affiliatewindow

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
	newPurchase.Head.PurchaseID = transaction.ID
	newPurchase.Head.CreatedAt = stamps.New().String()

	newPurchase.Body.ClickID = getClickRef(transaction)
	newPurchase.Body.Status = statusOf(transaction)

	newPurchase.Body.Currency = transaction.AwinSaleAmount.Currency

	timestamp, flawError := stamps.ParseStamp(transaction.TransactionDate, stamps.YYYYhMMhDDTHHcMMcSS)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.Date = timestamp.String()

	saleAmount, flawError := currency.NewAmountFromString(transaction.AwinSaleAmount.Amount, transaction.AwinSaleAmount.Currency)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.SaleAmount = saleAmount.Value.String()

	commissionAmount, flawError := currency.NewAmountFromString(transaction.AwinCommissionAmount.Amount, transaction.AwinCommissionAmount.Currency)
	if flawError != nil {
		return nil, flawError
	}
	newPurchase.Body.CommissionAmount = commissionAmount.Value.String()

	// For AWIN - we need to set sale and commission to -ve if it is declined (IN-582 and INTICKETS-250)
	if newPurchase.Body.Status == networkaccountpurchase.DECLINED {
		newPurchase.Body.SaleAmount = "-" + newPurchase.Body.SaleAmount
		newPurchase.Body.CommissionAmount = "-" + newPurchase.Body.CommissionAmount
	}

	newPurchase.Body.OrderID = transaction.OrderRef

	if newPurchase.Body.OrderID == "" {
		newPurchase.Body.OrderID = newPurchase.ID()
	}

	newPurchase.Body.Variant = Variant
	newPurchase.Body.Variations = transaction

	return newPurchase, nil
}

func getClickRef(transaction transaction) string {
	if transaction.ClickRefs.ClickRef != "" {
		return transaction.ClickRefs.ClickRef
	}

	if transaction.ClickRefs.ClickRef2 != "" {
		return transaction.ClickRefs.ClickRef2
	}

	if transaction.ClickRefs.ClickRef3 != "" {
		return transaction.ClickRefs.ClickRef3
	}

	if transaction.ClickRefs.ClickRef4 != "" {
		return transaction.ClickRefs.ClickRef4
	}

	if transaction.ClickRefs.ClickRef5 != "" {
		return transaction.ClickRefs.ClickRef5
	}

	if transaction.ClickRefs.ClickRef6 != "" {
		return transaction.ClickRefs.ClickRef6
	}

	return ""
}
