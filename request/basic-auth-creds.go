package request

func (r *Request) BasicAuthCredentials() (string, string, bool) {
	return r.BasicAuth()
}
