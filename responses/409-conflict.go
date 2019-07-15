package responses

import "net/http"

func (response *Response) Conflict() {
	response.HTTPStatusCode = http.StatusConflict
}
