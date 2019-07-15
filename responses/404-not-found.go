package responses

import "net/http"

func (response *Response) NotFound() {
	response.HTTPStatusCode = http.StatusNotFound
}
