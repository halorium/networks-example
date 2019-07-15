package server

import (
	"net/http"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/go/gc-metrics-logger"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/profiler"
	"github.com/halorium/networks-example/router"
)

// ListenAndServe starts an HTTP(S) server
func ListenAndServe(router *router.Router) {
	profiler.ListenAndServe()

	gcmetricslogger.Start()

	router.GET("/status", getStatus)

	if envvars.Mocked {
		go func() {
			logger.Debug("server-listen-and-serve", nil)

			err := http.ListenAndServe(":80", router)

			if err != nil {
				logger.Panic(
					"server-listen-and-serve",
					flaw.From(err).Wrap("cannot ListenAndServe"),
				)
			}
		}()
	} else {
		go func() {
			logger.Debug("server-listen-and-serve-tls", router)

			err := http.ListenAndServeTLS(":443", "/tls/files/certificate-chain.pem", "/tls/files/private-key.pem", router)

			if err != nil {
				logger.Panic(
					"server-listen-and-serve-tls",
					flaw.From(err).Wrap("cannot ListenAndServe"),
				)
			}
		}()
	}

	select {}
}
