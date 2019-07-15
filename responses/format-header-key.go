package responses

import (
	"regexp"
	"strings"
)

var reg = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func FormatHeaderKey(prefix string, header string) string {
	var a []string
	a = append(a, prefix)
	for _, sub := range reg.FindAllStringSubmatch(header, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	//return strings.ToLower(strings.Join(a, "_"))
	return strings.Join(a, "-")
}
