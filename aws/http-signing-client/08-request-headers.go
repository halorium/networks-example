package awshttpsigningclient

func requestHeaders(state *State) {
	state.Request.Header.Set(
		"Host",
		state.Request.Host,
	)

	state.Request.Header.Set(
		"X-Amz-Content-Sha256",
		state.RequestBodySHA256Hex,
	)

	state.Request.Header.Set(
		"X-Amz-Date",
		state.XAmzDate,
	)

	if state.Creds.Token != "" {
		state.Request.Header.Set(
			"X-Amz-Security-Token",
			state.Creds.Token,
		)
	}

	sortedLowercaseHeaderNames(state)
}
