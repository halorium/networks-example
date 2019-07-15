package router

import (
	"net/http"

	"github.com/halorium/httprouter"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/request"
	"github.com/halorium/networks-example/state"
)

func Endpoint(tag string, endpoint func(*state.State)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		state := state.New()
		state.Request = request.NewRequest(w, r, &p)
		state.Response.Writer = w

		logger.Info("http-request", state.Request.Details())

		state.Nickname = tag

		endpoint(state)

		state.Response.Send()

		logger.Info("http-response", state)
	}
}
