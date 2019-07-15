package main

import (
	"github.com/halorium/networks-example/router"
	"github.com/halorium/networks-example/server"
)

func main() {
	rolesRouter := router.New()

	rolesRouter.PUT(
		"/roles/:role-id",
		router.Endpoint("put-role", putRole),
	)

	rolesRouter.GET(
		"/roles/:role-id",
		router.Endpoint("get-role", getRole),
	)

	rolesRouter.GET(
		"/roles/:role-id/history",
		router.Endpoint("get-role-history", getRoleHistory),
	)

	server.ListenAndServe(rolesRouter)
}
