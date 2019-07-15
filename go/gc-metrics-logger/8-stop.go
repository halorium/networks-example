package gcmetricslogger

import "time"

func (metrics *Metrics) Stop() {
	metrics.stopChannel <- true

	time.Sleep(time.Second)
}
