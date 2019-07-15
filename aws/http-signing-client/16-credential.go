package awshttpsigningclient

func credential(state *State) {
	state.Credential = state.Creds.AccessKeyID + "/" + state.Scope

	stringToSign(state)
}
