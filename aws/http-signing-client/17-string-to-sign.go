package awshttpsigningclient

import "strings"

func stringToSign(state *State) {
	state.StringToSign = strings.Join(
		[]string{
			ALGORITHM,
			state.XAmzDate,
			state.Scope,
			state.CanonicalRequestSHA256Hex,
		},
		"\n",
	)

	authorizationHeader(state)
}
