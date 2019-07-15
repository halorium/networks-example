package awshttpsigningclient

import (
	"net/http"

	"github.com/halorium/networks-example/http-client"
)

type Transport struct {
	State *State
}

func (t Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	t.State.Request = *request
	sign(t.State)

	httpClient := httpclient.FromPool()

	res, err := httpClient.Do(request)

	httpclient.ToPool(httpClient)

	return res, err
}
