package request

import "strings"

func (r *Request) BearerAuthToken() (string, bool) {
	auth := r.Header.Get("Authorization")

	if !strings.HasPrefix(auth, "Bearer ") {
		return "", false
	}

	return auth[7:], true
}
