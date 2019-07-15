package flaw

func (err *Error) Error() string {
	return err.MessageTrace[len(err.MessageTrace)-1].Message
}
