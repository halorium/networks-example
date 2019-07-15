package booking

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/networks/network-account-purchase"
)

func statusOf(transaction transaction) (string, *flaw.Error) {
	// Booking.com doesn't give a status with the data
	return networkaccountpurchase.OPEN, nil
}
