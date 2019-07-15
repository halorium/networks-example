package main

import (
	"io/ioutil"
	"net/http"

	"github.com/halorium/httprouter"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/router"
)

func main() {
	router := router.New()

	router.GET("/partner/login.html", getLogin)

	router.GET("/affiliate/reservations_details.html", getTransactions)

	panic(
		http.ListenAndServe(":80", router),
	)
}

func getLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(200)

	res, err := ioutil.ReadFile("login-response.html")

	if err != nil {
		logger.Panic("networks-booking-mock-api", flaw.From(err))
	}

	_, err = w.Write(res)

	if err != nil {
		logger.Panic("networks-booking-mock-api", flaw.From(err))
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=data.csv")
	w.WriteHeader(200)

	res, err := ioutil.ReadFile("response.csv")
	if err != nil {
		logger.Panic("networks-booking-mock-api", flaw.From(err))
	}

	_, err = w.Write(res)

	if err != nil {
		logger.Panic("networks-booking-mock-api", flaw.From(err))
	}
}
