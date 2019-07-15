package responses

import "net/http"

// Created HTTP response
func (response *Response) Created() {
	response.HTTPStatusCode = http.StatusCreated
}
