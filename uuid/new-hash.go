package uuid

import (
	"crypto/sha256"
	"fmt"
)

// NewHash ** Generates a hash ID from one or more strings **
func NewHash(inputs ...string) string {
	concatenated := ""

	for _, input := range inputs {
		concatenated += input
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(concatenated)))
}
