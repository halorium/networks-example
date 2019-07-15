package flaw

import "runtime"

func (err *Error) Wrap(message string) *Error {
	_, pathname, line, ok := runtime.Caller(1)

	if !ok {
		panic("not ok")
	}

	messageTraceStruct := messageTrace{
		Message:  message,
		Pathname: stripPathname(pathname),
		Line:     line,
	}

	err.MessageTrace = append([]messageTrace{messageTraceStruct}, err.MessageTrace...)

	return err
}
