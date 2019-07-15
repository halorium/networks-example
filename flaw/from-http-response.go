package flaw

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func FromHTTPResponse(res *http.Response, err error) (*http.Response, *Error) {
	if err != nil {
		return res, create(err.Error())
	}

	if !strings.HasPrefix(res.Status, "2") {
		bytes, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return res, create(err.Error())
		}

		return res, create(res.Status + "\n\n" + string(bytes))
	}

	return res, nil
}
