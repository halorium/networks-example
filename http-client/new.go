package httpclient

import (
	"net/http"
	"time"
)

func New() *Client {
	return &Client{
		&http.Client{
			Timeout: 60 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}
