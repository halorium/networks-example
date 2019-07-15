package awshttpsigningclient

import "strings"

func canonicalRequest(state *State) {
	state.CanonicalRequest = strings.Join(
		[]string{
			state.Request.Method,
			state.CanonicalURI,
			state.CanonicalQuery,
			state.CanonicalHeaders,
			state.SignedHeaders,
			state.RequestBodySHA256Hex,
		},
		"\n",
	)

	canonicalRequestSHA256Hex(state)
}
