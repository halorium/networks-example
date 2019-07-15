package main

import "github.com/halorium/networks-example/worker"

func main() {
	networkWorker := worker.New()

	networkWorker.Receive(
		"post-network-account-purchases-load",
		postNetworkAccountPurchasesLoad,
	)

	networkWorker.Receive(
		"put-purchase-in-postgres",
		putNetworkAccountPurchaseInPostgres,
	)

	networkWorker.Process()
}
