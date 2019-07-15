package flaw

func (err *Error) stackTrace() string {
	stackTraceString := ""

	for _, stackTrace := range err.StackTrace {
		stackTraceString += stackTrace.String() + "\n"
	}

	return stackTraceString
}
