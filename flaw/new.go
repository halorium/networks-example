package flaw

func New(message string) *Error {
	return create(message)
}
