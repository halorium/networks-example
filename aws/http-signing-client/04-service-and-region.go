package awshttpsigningclient

import "strings"

func serviceAndRegion(state *State) {
	parts := strings.Split(state.Request.Host, ".")

	state.Service = parts[len(parts)-3]
	state.Region = parts[len(parts)-4]

	canonicalQuery(state)
}
