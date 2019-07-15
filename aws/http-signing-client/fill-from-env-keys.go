package awshttpsigningclient

func fillFromEnvKeys(state *State) bool {
	if state.Creds.AccessKeyID == "" ||
		state.Creds.SecretAccessKey == "" ||
		state.CredsFromSTS {
		return false
	}

	return true
}
