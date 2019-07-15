package entities

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/http-client"
)

type Entity interface {
	Prefix() string
	ID() string
	Etag() string
	CreatedAt() string
	DBJSON() []byte
	BodyJSON() []byte
}

func Path(e Entity) string {
	return e.Prefix() + "/" + e.ID()
}

func PutToAPI(host string, e Entity) *flaw.Error {
	req, err := http.NewRequest(
		"PUT",
		host+e.Prefix()+"/"+url.PathEscape(e.ID()),
		bytes.NewReader(e.BodyJSON()),
	)

	if err != nil {
		return flaw.From(err).Wrap("cannot PutToAPI")
	}

	httpClient := httpclient.FromPool()

	_, flawError := flaw.FromHTTPResponse(
		httpClient.Do(req),
	)

	httpclient.ToPool(httpClient)

	if flawError != nil {
		return flawError.Wrap("cannot PutToAPI")
	}

	return nil
}
