package affilinet

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/http-client"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/stamps"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/try"
)

func fetchPage(state *state.State, apiToken string, paginator *paginator) *flaw.Error {
	logger.Debug("affilinet-fetch-page", state)

	url := envvars.AffilinetAPIHost + "/V2.0/PublisherStatistics.svc"

	startTime := state.NetworkAccountPurchasesLoad.StartTime.Format(stamps.YYYYhMMhDD)
	stopTime := state.NetworkAccountPurchasesLoad.StopTime.Format(stamps.YYYYhMMhDD)

	requestBuffer := bytes.NewBufferString(`
    <x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:svc="http://affilinet.framework.webservices/Svc" xmlns:pub="http://affilinet.framework.webservices/types/PublisherStatistics" xmlns:arr="http://schemas.microsoft.com/2003/10/Serialization/Arrays">
      <x:Header/>
      <x:Body>
        <svc:GetTransactionsRequest>
          <svc:CredentialToken>` + apiToken + `</svc:CredentialToken>
          <svc:PageSettings>
            <pub:CurrentPage>` + strconv.Itoa(paginator.Page) + `</pub:CurrentPage>
            <pub:PageSize>` + strconv.Itoa(paginator.Size) + `</pub:PageSize>
          </svc:PageSettings>
          <svc:TransactionQuery>
            <pub:EndDate>` + stopTime + `</pub:EndDate>
            <pub:StartDate>` + startTime + `</pub:StartDate>
            <pub:TransactionStatus>All</pub:TransactionStatus>
          </svc:TransactionQuery>
        </svc:GetTransactionsRequest>
      </x:Body>
    </x:Envelope>
  `)

	try := try.New()

	for {
		req, err := http.NewRequest("POST", url, requestBuffer)

		flawError := flaw.From(err)

		if flawError != nil {
			return flawError.Wrap("cannot fetchPage")
		}

		req.Header.Add("Content-Type", "text/xml; charset=utf-8")
		req.Header.Add("SOAPAction", "http://affilinet.framework.webservices/Svc/PublisherStatisticsContract/GetTransactions")

		res, flawError := httpclient.Make(req)

		if flawError != nil {
			if try.Again("affilinet-fetch-page-retry", flawError) {
				continue
			}

			return flawError.Wrap("cannot fetchPage")
		}

		state.VariantReadCloser = res.Body

		break
	}

	return deserializeXML(state)
}
