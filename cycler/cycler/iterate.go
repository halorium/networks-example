package cycler

import (
	"time"

	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/stamps"
)

func Iterate(cycle *Cycle) {
	cycle.StopTime = stamps.EndOfDay(stamps.New().PreviousDay())
	cycle.StartTime = stamps.BeginningOfDay(cycle.StopTime.OffsetByDaysPerRequest(cycle.DaysPerRequest))

	daysLeft := cycle.DaysToLoad

	for {
		flawError := postNetworkAccountPurchasesLoad(cycle)

		if flawError != nil {
			logger.Debug(cycle.Tag+"-fail", cycle)
		} else {
			logger.Debug(cycle.Tag+"-success", cycle)
		}

		daysLeft -= cycle.DaysPerRequest

		if daysLeft <= 0 {
			return
		}

		time.Sleep(cycle.Sleep)

		cycle.StopTime = stamps.EndOfDay(cycle.StartTime.PreviousDay())
		cycle.StartTime = stamps.BeginningOfDay(cycle.StopTime.OffsetByDaysPerRequest(cycle.DaysPerRequest))
	}
}
