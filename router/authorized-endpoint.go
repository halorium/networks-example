package router

import (
	"net/http"

	"github.com/halorium/httprouter"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/json-web-token"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/request"
	"github.com/halorium/networks-example/state"
)

func AuthorizedEndpoint(tag string, endpoint func(*state.State)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		state := state.New()
		state.Request = request.NewRequest(w, r, &p)
		state.Response.Writer = w

		logger.Info("http-request", state.Request.Details())

		state.Nickname = tag

		signedToken, ok := state.Request.BearerAuthToken()

		if !ok {
			state.FlawError = flaw.New("cannot AuthorizedEndpoint")
			state.Response.Head = nil
			state.Response.Body = nil
			state.Response.Unauthorized()
			state.Response.Send()
			return
		}

		JSONWebToken, err := jsonwebtoken.Parse(signedToken)

		if err != nil {
			state.FlawError = flaw.From(err).Wrap("cannot AuthorizedEndpoint")
			state.Response.Head = nil
			state.Response.Body = nil
			state.Response.Unauthorized()
			state.Response.Send()
			return
		}

		state.JSONWebToken = JSONWebToken

		endpoint(state)

		state.Response.Send()

		logger.Info("http-response", state)
	}
}
