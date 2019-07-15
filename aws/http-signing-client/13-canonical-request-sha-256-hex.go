package awshttpsigningclient

func canonicalRequestSHA256Hex(state *State) {
	state.CanonicalRequestSHA256Hex = sha256Hex([]byte(state.CanonicalRequest))

	signingKey(state)
}
