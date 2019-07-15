package awshttpsigningclient

// Sign signs a request with Signed Signature Version 4
func sign(state *State) {
	authentication(state)
}
