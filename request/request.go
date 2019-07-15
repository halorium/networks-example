package request

import (
	"net/http"

	"github.com/halorium/httprouter"
)

type Request struct {
	*http.Request
	Params *httprouter.Params `json:"params,omitempty"`
}

func NewRequest(w http.ResponseWriter, r *http.Request, p *httprouter.Params) *Request {
	r.Body = http.MaxBytesReader(w, r.Body, 1e6) // 1 MB maximum input body

	return &Request{r, p}
}
