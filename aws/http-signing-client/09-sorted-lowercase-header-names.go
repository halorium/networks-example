package awshttpsigningclient

import (
	"sort"
	"strings"
)

func sortedLowercaseHeaderNames(state *State) {
	sortedLowercaseHeaderNames := make([]string, 0, 24)

	for name := range state.Request.Header {
		lowerCaseName := strings.ToLower(name)

		sortedLowercaseHeaderNames = append(sortedLowercaseHeaderNames, lowerCaseName)
	}

	sort.Strings(sortedLowercaseHeaderNames)

	state.SortedLowercaseHeaderNames = sortedLowercaseHeaderNames

	signedHeaders(state)
}
