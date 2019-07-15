package affilinet

import "github.com/halorium/networks-example/networks/network-account-purchase"

func statusOf(transaction transaction) string {
	switch transaction.TransactionStatus {
	case "Open", "Confirmed":
		return networkaccountpurchase.OPEN
	case "Cancelled":
		return networkaccountpurchase.DECLINED
	default:
		return networkaccountpurchase.OPEN
	}
}
