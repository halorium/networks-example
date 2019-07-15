package flaw

import "runtime"

func create(message string) *Error {
	_, pathname, line, ok := runtime.Caller(2)

	if !ok {
		panic("not ok")
	}

	err := &Error{
		MessageTrace: []messageTrace{
			{
				Message:  message,
				Pathname: stripPathname(pathname),
				Line:     line,
			},
		},
		StackTrace: getFrames(),
	}

	return err
}
