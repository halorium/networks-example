package responses

import "net/http"

func (response *Response) NotImplemented() {
	response.HTTPStatusCode = http.StatusNotImplemented
}
