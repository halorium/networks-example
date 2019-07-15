package main

import (
	"io/ioutil"
	"net/http"

	"github.com/halorium/httprouter"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func main() {
	router := httprouter.New()

	router.GET("/publishers/68727/transactions/", get)

	panic(
		http.ListenAndServe(":80", router),
	)
}

func get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(200)

	res, err := ioutil.ReadFile("response.json")

	if err != nil {
		logger.Panic("networks-affiliate-window-mock-api", flaw.From(err))
	}

	_, err = w.Write(res)

	if err != nil {
		logger.Panic("networks-affiliate-window-mock-api", flaw.From(err))
	}
}
