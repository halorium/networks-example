package cycler

import (
	"time"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/panic-handler"
	"github.com/halorium/networks-example/sleep"
)

func (cycle *Cycle) Run() {
	logger.Info("cycler-run", cycle)

	defer panichandler.Log("cycler-run")

	duration, err := time.ParseDuration(cycle.Period)

	if err != nil {
		logger.Critical(cycle.Tag, flaw.From(err).Wrap("cannot Run"))

		return
	}

	sleep.ForAwhile(duration)

	for {
		cycle.Action(cycle)
		time.Sleep(duration)
	}
}
