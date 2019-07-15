package cycler

import (
	"strconv"
	"time"

	"github.com/halorium/networks-example/flaw"
)

func CalculateSleep(rateQuantity string, rateInterval string) (time.Duration, *flaw.Error) {
	var sleepDuration time.Duration

	interval, err := time.ParseDuration(rateInterval)

	if err != nil {
		return sleepDuration, flaw.From(err).Wrap("cannot CalculateSleep")
	}

	quantity, err := strconv.ParseFloat(rateQuantity, 64)

	if err != nil {
		return sleepDuration, flaw.From(err).Wrap("cannot CalculateSleep")
	}

	sleepNanoseconds := float64(interval) / quantity

	sleepDuration = time.Duration(sleepNanoseconds) + time.Second

	return sleepDuration, nil
}
