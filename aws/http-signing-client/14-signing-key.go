package awshttpsigningclient

func signingKey(state *State) {
	queDate := hmacSHA256Binary([]byte("AWS4"+state.Creds.SecretAccessKey), state.RequestDate)
	queRegion := hmacSHA256Binary(queDate, state.Region)
	queService := hmacSHA256Binary(queRegion, state.Service)

	state.SigningKey = hmacSHA256Binary(queService, SIGNATUREVERSION)

	scope(state)
}
