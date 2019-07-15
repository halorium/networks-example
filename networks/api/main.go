package main

import (
	"github.com/halorium/networks-example/router"
	"github.com/halorium/networks-example/server"
)

func main() {
	networkRouter := router.New()

	networkRouter.GET(
		"/networks",
		router.Endpoint("get-networks", getNetworks),
	)

	networkRouter.GET(
		"/networks/:network-id",
		router.Endpoint("get-network", getNetwork),
	)

	networkRouter.GET(
		"/networks/:network-id/history",
		router.Endpoint("get-network-history", getNetworkHistory),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts",
		router.Endpoint("get-network-accounts", getNetworkAccounts),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts/:network-account-id",
		router.Endpoint("get-network-account", getNetworkAccount),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts/:network-account-id/history",
		router.Endpoint("get-network-account-history", getNetworkAccountHistory),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts/:network-account-id/purchases",
		router.Endpoint("get-network-account-purchases", getNetworkAccountPurchases),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts/:network-account-id/purchase-updates",
		router.Endpoint("get-network-account-purchase-updates", getNetworkAccountPurchaseUpdates),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts/:network-account-id/purchases/:purchase-id",
		router.Endpoint("get-network-account-purchase", getNetworkAccountPurchase),
	)

	networkRouter.GET(
		"/networks/:network-id/accounts/:network-account-id/purchases/:purchase-id/history",
		router.Endpoint("get-network-account-purchase-history", getNetworkAccountPurchaseHistory),
	)

	// DELETE endpoints need to be secured - waiting for TLS encryption

	networkRouter.DELETE(
		"/networks/:network-id",
		router.Endpoint("delete-network", deleteNetwork),
	)

	networkRouter.DELETE(
		"/networks/:network-id/accounts",
		router.Endpoint("delete-network-accounts", deleteNetworkAccounts),
	)

	networkRouter.DELETE(
		"/networks/:network-id/accounts/:network-account-id",
		router.Endpoint("delete-network-account", deleteNetworkAccount),
	)

	networkRouter.DELETE(
		"/networks/:network-id/accounts/:network-account-id/purchases",
		router.Endpoint("delete-network-account", deleteNetworkAccountPurchases),
	)

	networkRouter.POST(
		"/networks/:network-id/accounts/:network-account-id/purchases/load",
		router.Endpoint("post-network-account-purchases-load", postNetworkAccountPurchasesLoad),
	)

	networkRouter.PUT(
		"/networks/:network-id",
		router.Endpoint("put-network", putNetwork),
	)

	networkRouter.PUT(
		"/networks/:network-id/accounts/:network-account-id",
		router.Endpoint("put-network-account", putNetworkAccount),
	)

	networkRouter.PUT(
		"/networks/:network-id/accounts/:network-account-id/purchases/:purchase-id",
		router.Endpoint("put-network-account-purchase", putNetworkAccountPurchase),
	)

	server.ListenAndServe(networkRouter)
}
