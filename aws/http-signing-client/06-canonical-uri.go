package awshttpsigningclient

import "strings"

func canonicalURI(state *State) {
	state.CanonicalURI = state.Request.URL.EscapedPath()

	// This is due to a bug / strange behaviour in AWS / Elasticsearch which
	// requires @ escaped even though URIs don't require it
	state.CanonicalURI = strings.Replace(state.CanonicalURI, "@", "%40", -1)

	// This is due to a bug / strange behaviour in AWS / Elasticsearch which
	// requires %28 and %29 (parenthesis) to be double-escaped
	state.CanonicalURI = strings.Replace(state.CanonicalURI, "%28", "%2528", -1)
	state.CanonicalURI = strings.Replace(state.CanonicalURI, "%29", "%2529", -1)

	// This is due to a bug / strange behaviour in AWS / Elasticsearch which
	// decodes %2F in the path to / and thus changes the path so it must be
	// double encoded
	state.CanonicalURI = strings.Replace(state.CanonicalURI, "%2F", "%252F", -1)

	requestBodySHA256Hex(state)
}
