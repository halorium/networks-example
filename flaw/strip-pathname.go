package flaw

import "regexp"

func stripPathname(pathname string) string {
	rgx := regexp.MustCompile(".+/halorium/networks-example/(.+)")

	pathMatches := rgx.FindAllStringSubmatch(pathname, -1)

	if pathMatches == nil {
		return pathname
	}

	return pathMatches[0][1]
}
