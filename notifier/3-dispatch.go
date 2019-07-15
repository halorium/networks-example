package notifier

import "github.com/halorium/networks-example/logger"

func dispatch(state *State) {
	flawError := state.Handler(state.Message)

	if flawError != nil {
		state.FlawError = flawError

		logger.Critical("notifer-dispatch-failure", state)

		return
	}

	logger.Info("notifier-dispatch-success", state)

	state.Message.Delete()
}
