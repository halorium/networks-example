package server

import (
	"net/http"

	"github.com/halorium/httprouter"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/request"
	"github.com/halorium/networks-example/state"
)

func getStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	state := state.New()
	state.Request = request.NewRequest(w, r, &p)
	state.Response.Writer = w

	logger.Debug("status-request", state.Request.Details())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	logger.Info("status-response", state)
}
