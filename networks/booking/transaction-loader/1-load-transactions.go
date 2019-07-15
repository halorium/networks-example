package booking

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/stamps"
	"github.com/halorium/networks-example/state"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

func LoadTransactions(state *state.State) *flaw.Error {
	logger.Debug("booking-load-transactions", state)

	headlessBrowser := surf.NewBrowser()
	headlessBrowser.SetAttribute(browser.FollowRedirects, false)

	loginURL := envvars.BookingAPIHost + "/partner/login.html"

	username := state.NetworkAccount.Body.Variations.Username
	password := state.NetworkAccount.Body.Variations.Password

	err := headlessBrowser.Open(loginURL)

	if err != nil {
		return flaw.From(err).Wrap("cannot LoadTransactions")
	}

	loginForm, err := headlessBrowser.Form("form")

	if err != nil {
		return flaw.From(err).Wrap("cannot LoadTransactions")
	}

	err = loginForm.Input("loginname", username)

	if err != nil {
		return flaw.From(err).Wrap("cannot LoadTransactions")
	}

	err = loginForm.Input("password", password)

	if err != nil {
		return flaw.From(err).Wrap("cannot LoadTransactions")
	}

	err = loginForm.Submit()

	var session string
	sessionRegexp := regexp.MustCompile(`ses=([a-zA-Z0-9_]+)`)

	if err != nil {
		if strings.Contains(err.Error(), "ses=") {
			match := sessionRegexp.FindStringSubmatch(err.Error())

			if len(match) < 2 {
				return flaw.New("authentication failure").Wrap("cannot LoadTransactions")
			}

			session = match[1]
		} else {
			return flaw.From(err).Wrap("cannot LoadTransactions")
		}
	}

	startTime := state.NetworkAccountPurchasesLoad.StartTime.Format(stamps.YYYYhMMhDD)
	affiliateID := state.NetworkAccount.Body.Variations.AdditionalID

	transactionsURL := envvars.BookingAPIHost +
		"/affiliate/reservations_details.html?" +
		"&affiliate_id=" + affiliateID +
		"&wide=1&date_field=created&interval=day" +
		"&start_date=" + startTime +
		"&count=0&format=csv" +
		"&ses=" + session

	err = headlessBrowser.Open(transactionsURL)

	if err != nil {
		return flaw.From(err).Wrap("cannot LoadTransactions")
	}

	bodyText := headlessBrowser.Dom().Find("body").Text()

	reader := bytes.NewReader([]byte(bodyText))

	readCloser := ioutil.NopCloser(reader)

	state.VariantReadCloser = readCloser

	return deserializeCSV(state)
}
