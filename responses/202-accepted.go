package responses

import "net/http"

func (response *Response) Accepted() {
	response.HTTPStatusCode = http.StatusAccepted
}
