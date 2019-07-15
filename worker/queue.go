package worker

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/notifier"
)

type Queue struct {
	Name    string
	Handler func(*notifier.Message) *flaw.Error
	URL     string
}
