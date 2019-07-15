package awshttpsigningclient

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func requestBodySHA256Hex(state *State) {
	var duplicateBody io.ReadCloser
	var err error

	if state.Request.Body == nil {
		duplicateBody = ioutil.NopCloser(strings.NewReader(""))
	} else {
		duplicateBody, err = state.Request.GetBody()

		if err != nil {
			state.FlawError = flaw.From(err).Wrap("cannot requestBodySHA256Hex")
			logger.Panic("request-body-sha-256-hex-get-body", state)
		}
	}

	defer duplicateBody.Close()

	bytes, err := ioutil.ReadAll(duplicateBody)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot requestBodySHA256Hex")
		logger.Panic("request-body-sha-256-hex-read-all", state)
	}

	state.RequestBodySHA256Hex = sha256Hex(bytes)

	requestHeaders(state)
}
