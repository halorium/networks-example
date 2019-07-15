package affilinet

import (
	"bytes"
	"net/http"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/http-client"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/try"
)

func authenticate(state *state.State) *flaw.Error {
	logger.Debug("affilinet-authenticate", state)

	url := envvars.AffilinetAPIHost + "/V2.0/Logon.svc"

	var apiToken string

	username := state.NetworkAccount.Body.Variations.Username
	password := state.NetworkAccount.Body.Variations.Password

	requestBuffer := bytes.NewBufferString(`
    <x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:svc="http://affilinet.framework.webservices/Svc" xmlns:typ="http://affilinet.framework.webservices/types">
      <x:Header/>
      <x:Body>
        <svc:LogonRequestMsg>
          <typ:Username>` + username + `</typ:Username>
          <typ:Password>` + password + `</typ:Password>
          <typ:WebServiceType>Publisher</typ:WebServiceType>
        </svc:LogonRequestMsg>
      </x:Body>
    </x:Envelope>
  `)

	try := try.New()

	for {
		req, err := http.NewRequest("POST", url, requestBuffer)

		flawError := flaw.From(err)

		if flawError != nil {
			return flawError.Wrap("cannot authenticate")
		}

		req.Header.Add("Content-Type", "text/xml; charset=utf-8")
		req.Header.Add("SOAPAction", "http://affilinet.framework.webservices/Svc/ServiceContract1/Logon")

		res, flawError := httpclient.Make(req)

		if flawError != nil {
			if try.Again("affilinet-authenticate-retry", flawError) {
				continue
			}

			return flawError.Wrap("cannot authenticate")
		}

		var auth authentication

		flawError = serializers.XMLDeserializer(res.Body, &auth)

		if flawError != nil {
			return flawError.Wrap("cannot authenticate")
		}

		apiToken = auth.CredentialToken

		break
	}

	return paginate(state, apiToken)
}
