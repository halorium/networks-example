package responses

import "net/http"

func (response *Response) ServiceUnavailable() {
	response.HTTPStatusCode = http.StatusServiceUnavailable
}
