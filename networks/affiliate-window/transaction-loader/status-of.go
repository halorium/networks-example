package affiliatewindow

import "github.com/halorium/networks-example/networks/network-account-purchase"

// Affiliate Window has the following Status values: pending, confirmed, declined
// Affiliate Window also has the PaidToPublisher boolean value.

func statusOf(transaction transaction) string {
	if transaction.PaidToPublisher == "true" {
		return networkaccountpurchase.PAID
	}

	switch transaction.CommissionStatus {
	case "confirmed", "pending":
		return networkaccountpurchase.OPEN
	case "declined":
		return networkaccountpurchase.DECLINED
	default:
		return networkaccountpurchase.OPEN
	}
}
