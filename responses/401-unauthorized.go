package responses

import "net/http"

func (response *Response) Unauthorized() {
	response.HTTPStatusCode = http.StatusUnauthorized
}
