package awshttpsigningclient

import "time"

func requestTimes(state *State) {
	state.RequestTime = time.Now().UTC()
	state.RequestDate = state.RequestTime.Format("20060102")
	state.XAmzDate = state.RequestTime.Format("20060102T150405Z")

	serviceAndRegion(state)
}
