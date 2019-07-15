package affiliatewindow

import (
	"net/http"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/http-client"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/try"
)

func LoadTransactions(state *state.State) *flaw.Error {
	logger.Debug("affiliate-window-load-transactions", state)

	additionalID := state.NetworkAccount.Body.Variations.AdditionalID
	token := state.NetworkAccount.Body.Variations.Token

	url := envvars.AffiliateWindowAPIHost + "/publishers/" + additionalID + "/transactions/"

	startTime := state.NetworkAccountPurchasesLoad.StartTime.String()
	stopTime := state.NetworkAccountPurchasesLoad.StopTime.String()

	try := try.New()

	for {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return flaw.From(err).Wrap("cannot LoadTransactions")
		}

		query := req.URL.Query()
		query.Add("startDate", startTime)
		query.Add("endDate", stopTime)
		req.URL.RawQuery = query.Encode()

		req.Header.Add("Authorization", "Bearer "+token)

		res, flawError := httpclient.Make(req)

		if flawError != nil {
			if try.Again("affiliate-window-load-transactions-retry", flawError) {
				continue
			}

			return flawError.Wrap("cannot LoadTransactions")
		}

		state.VariantReadCloser = res.Body

		break
	}

	return deserializeJSON(state)
}
