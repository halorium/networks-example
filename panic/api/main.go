package main

import (
	"github.com/halorium/networks-example/router"
	"github.com/halorium/networks-example/server"
)

func main() {
	panicRouter := router.New()

	panicRouter.GET(
		"/panic",
		router.Endpoint("get-panic", getPanic),
	)

	server.ListenAndServe(panicRouter)
}
