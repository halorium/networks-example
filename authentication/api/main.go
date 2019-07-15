package main

import (
	"github.com/halorium/networks-example/router"
	"github.com/halorium/networks-example/server"
)

func main() {
	authenticationRouter := router.New()

	authenticationRouter.POST(
		"/authentication",
		router.Endpoint("post-authentication", postAuthentication),
	)

	authenticationRouter.POST(
		"/authentication/validate",
		router.AuthorizedEndpoint("post-authentication-validate", postAuthenticationValidate),
	)

	server.ListenAndServe(authenticationRouter)
}
