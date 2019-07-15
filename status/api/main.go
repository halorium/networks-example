package main

import (
	"github.com/halorium/networks-example/router"
	"github.com/halorium/networks-example/server"
)

func main() {
	statusRouter := router.New()

	server.ListenAndServe(statusRouter)
}
