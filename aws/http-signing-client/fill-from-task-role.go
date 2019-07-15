package awshttpsigningclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/try"
)

func fillFromTaskRole(state *State) bool {
	logger.Debug("aws-signing-fill-from-task-role-expired", state)

	if time.Now().Before(state.RefreshAvailableAt) {
		return true
	}

	uri := "http://169.254.170.2" + envvars.AwsContainerCredentialsRelativeURI

	var res *http.Response
	var flawError *flaw.Error

	try := try.New()

	for {
		res, flawError = flaw.FromHTTPResponse(
			http.Get(uri),
		)

		if flawError != nil {
			if try.Again("aws-signing-fill-from-task-role-retry", flawError) {
				continue
			}

			state.FlawError = flawError.Wrap("cannot fillFromTaskRole")

			return false
		}

		break
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot fillFromTaskRole")
		logger.Warn("aws-signing-fill-from-task-role-read-all", state)
		return false
	}

	defer res.Body.Close()

	err = json.Unmarshal(body, state.Creds)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot fillFromTaskRole")
		logger.Warn("aws-signing-fill-from-task-role-unmarshal", state)
		return false
	}

	expiration, err := time.Parse(time.RFC3339, state.Creds.Expiration)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot fillFromTaskRole")
		logger.Warn("aws-signing-fill-from-task-role-parse", state)
		return false
	}

	state.RefreshAvailableAt = expiration.Add(-4 * time.Minute)
	state.CredsFromSTS = true

	return true
}
