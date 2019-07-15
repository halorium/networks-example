package responses

import "net/http"

// Ok HTTP response
func (response *Response) Ok() {
	response.HTTPStatusCode = http.StatusOK
}
