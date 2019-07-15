package main

import (
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func postAuthenticationValidate(state *state.State) {
	logger.Debug("post-authentication-validate", state)

	state.Response.Head = nil
	state.Response.Body = "valid"
	state.Response.Ok()
}
