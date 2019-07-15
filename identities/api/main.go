package main

import (
	"github.com/halorium/networks-example/router"
	"github.com/halorium/networks-example/server"
)

func main() {
	identitiesRouter := router.New()

	identitiesRouter.POST(
		"/identities",
		router.Endpoint("post-identity", postIdentity),
	)

	identitiesRouter.GET(
		"/identities/:identity-id",
		router.Endpoint("get-identity", getIdentity),
	)

	identitiesRouter.GET(
		"/identities/:identity-id/history",
		router.Endpoint("get-identity-history", getIdentityHistory),
	)

	identitiesRouter.PUT(
		"/identities/:identity-id/role-set",
		router.Endpoint("put-identity-role-set", putIdentityRoleSet),
	)

	server.ListenAndServe(identitiesRouter)
}
