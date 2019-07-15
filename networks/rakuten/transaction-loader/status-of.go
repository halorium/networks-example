package rakuten

import "github.com/halorium/networks-example/networks/network-account-purchase"

// Rakuten doesn't send status with a transaction
func statusOf(transaction transaction) string {
	return networkaccountpurchase.OPEN
}
