package flaw

type Error struct {
	MessageTrace []messageTrace `json:"message-trace"`
	StackTrace   []frame        `json:"stack-trace"`
}
