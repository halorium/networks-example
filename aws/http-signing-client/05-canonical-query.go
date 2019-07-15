package awshttpsigningclient

import "strings"

func canonicalQuery(state *State) {
	query := state.Request.URL.Query().Encode()

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in
	// the original query string has been percent escaped so all '+' chars that
	// are left were originally spaces.

	state.CanonicalQuery = strings.Replace(query, "+", "%20", -1)

	canonicalURI(state)
}
