package flaw

func From(err error) *Error {
	if err == nil {
		return nil
	}

	return create(err.Error())
}
