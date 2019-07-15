package awshttpsigningclient

import (
	"crypto/sha256"
	"fmt"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func sha256Hex(content []byte) string {
	h := sha256.New()
	_, err := h.Write(content)

	if err != nil {
		flawError := flaw.From(err).Wrap("cannot sha256Hex")
		logger.Panic("sha2256-hex", flawError)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
