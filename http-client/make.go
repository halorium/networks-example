package httpclient

import (
	"net/http"

	"github.com/halorium/networks-example/flaw"
)

func Make(req *http.Request) (*http.Response, *flaw.Error) {
	client := FromPool()

	res, flawError := flaw.FromHTTPResponse(
		client.Do(req),
	)

	ToPool(client)

	if flawError != nil {
		return nil, flawError
	}

	return res, nil
}
