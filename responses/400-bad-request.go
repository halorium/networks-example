package responses

import "net/http"

func (response *Response) BadRequest() {
	response.HTTPStatusCode = http.StatusBadRequest
}
