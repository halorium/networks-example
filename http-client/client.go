package httpclient

import "net/http"

type Client struct {
	*http.Client
}
