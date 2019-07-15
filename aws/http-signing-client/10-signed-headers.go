package awshttpsigningclient

import "strings"

func signedHeaders(state *State) {
	state.SignedHeaders = strings.Join(state.SortedLowercaseHeaderNames, ";")

	canonicalHeaders(state)
}
