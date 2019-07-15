package flaw

type messageTrace struct {
	Message  string `json:"message"`
	Pathname string `json:"pathname"`
	Line     int    `json:"line"`
}
