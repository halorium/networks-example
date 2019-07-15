package request

type Details struct {
	RequestMethod string `json:"request-method"`
	RequestURL    string `json:"request-url"`
}

func (req *Request) Details() *Details {
	return &Details{
		RequestMethod: req.Method,
		RequestURL:    req.URL.String(),
	}
}
