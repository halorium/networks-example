package awshttpsigningclient

import "strings"

func scope(state *State) {
	state.Scope = strings.Join(
		[]string{
			state.RequestDate,
			state.Region,
			state.Service,
			SIGNATUREVERSION,
		},
		"/",
	)

	credential(state)
}
