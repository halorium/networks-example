package awshttpsigningclient

import "strings"

func canonicalHeaders(state *State) {
	canonicalHeaders := make([]string, 0, 24)

	for _, lowercaseName := range state.SortedLowercaseHeaderNames {
		value := strings.TrimSpace(state.Request.Header.Get(lowercaseName))

		if lowercaseName == "host" {
			parts := strings.Split(lowercaseName, ":")

			if len(parts) > 1 {
				port := parts[1]

				if port == "80" || port == "443" {
					value = parts[0]
				}
			}
		}

		condensedValue := state.SpaceRunRegexp.ReplaceAllString(value, " ")

		canonicalHeaders = append(canonicalHeaders, lowercaseName+":"+condensedValue)
	}

	state.CanonicalHeaders = strings.Join(canonicalHeaders, "\n") + "\n"

	canonicalRequest(state)
}
