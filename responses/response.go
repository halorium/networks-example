package responses

import (
	"net/http"

	"github.com/halorium/networks-example/serializers"
)

type Response struct {
	Head           interface{}            `json:"head,omitempty"`
	Body           interface{}            `json:"body,omitempty"`
	Writer         http.ResponseWriter    `json:"-"`
	State          interface{}            `json:"state,omitempty"`
	HTTPStatusCode int                    `json:"http-status-code"`
	Serializer     serializers.Serializer `json:"-"`
	Sent           bool                   `json:"sent"`
}

func NewResponse() *Response {
	return &Response{}
}
