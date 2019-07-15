package profiler

import (
	"net/http"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"

	// stdlib made me do it!
	_ "net/http/pprof"
)

// ListenAndServe starts an HTTP(S) net/http/pprof profiler server
func ListenAndServe() {
	if envvars.Mocked {
		go func() {
			logger.Debug("profiler-listen-and-serve", nil)

			err := http.ListenAndServe(":81", nil)

			if err != nil {
				logger.Panic(
					"profiler-listen-and-serve",
					flaw.From(err).Wrap("cannot ListenAndServe"),
				)
			}
		}()
	} else {
		go func() {
			logger.Debug("profiler-listen-and-serve-tls", nil)

			err := http.ListenAndServeTLS(":444", "/tls/files/certificate-chain.pem", "/tls/files/private-key.pem", nil)

			if err != nil {
				logger.Panic(
					"profiler-listen-and-serve-tls",
					flaw.From(err).Wrap("cannot ListenAndServe"),
				)
			}
		}()
	}
}
