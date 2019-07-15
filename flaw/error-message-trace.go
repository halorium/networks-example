package flaw

import "fmt"

func (err *Error) messageTrace() string {
	messageTraceString := ""

	for _, messageTrace := range err.MessageTrace {
		messageTraceString += fmt.Sprintf(
			"%s (%s:%d)\n",
			messageTrace.Message,
			messageTrace.Pathname,
			messageTrace.Line,
		)
	}

	return messageTraceString
}
