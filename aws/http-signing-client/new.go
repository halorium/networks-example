package awshttpsigningclient

import (
	"net/http"

	"github.com/halorium/networks-example/http-client"
)

func New() (*httpclient.Client, *Credentials) {
	state := NewState()

	transport := Transport{
		State: state,
	}

	httpClient := &http.Client{
		Transport: http.RoundTripper(transport),
	}

	client := &httpclient.Client{
		Client: httpClient,
	}

	return client, state.Creds
}
