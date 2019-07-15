package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account-purchase-load"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func postNetworkAccountPurchasesLoad(state *state.State) {
	logger.Debug("post-network-account-purchases-load", state)

	// validate network
	getNetwork(state)

	if state.FlawError != nil {
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.BadRequest()
		return
	}

	// validate network-account
	getNetworkAccount(state)

	if state.FlawError != nil {
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.BadRequest()
		return
	}

	state.NetworkAccountPurchasesLoad = networkaccountpurchaseload.NewEntity()

	flawError := serializers.JSONDeserializer(state.Request.Body, state.NetworkAccountPurchasesLoad)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot postNetworkAccountPurchasesLoad")

		if strings.Contains(flawError.Error(), "parsing time") &&
			strings.Contains(flawError.Error(), "cannot parse") {
			state.Response.Head = nil
			state.Response.Body = validation.NewFrom("time fields must be formatted as 2006-01-02T15:04:05.000000Z")
			state.Response.BadRequest()
			return
		}

		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := networkaccountpurchaseload.Validate(state.NetworkAccountPurchasesLoad)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	days, err := strconv.Atoi(state.Network.Body.Variations.DaysPerRequest)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot postNetworkAccountPurchasesLoad")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("cannot parse days-per-request")
		state.Response.BadRequest()
		return
	}

	requestedDays := int(state.NetworkAccountPurchasesLoad.StopTime.Sub(*state.NetworkAccountPurchasesLoad.StartTime.Time).Hours() / 24)

	if requestedDays > days || (requestedDays == 1 && days == 1) {
		message := fmt.Sprintf("range must be %v day(s) or less", days)
		state.FlawError = flaw.New(message)
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom(message)
		state.Response.BadRequest()
		return
	}

	n := notifier.FromPool()

	flawError = n.Notify("post-network-account-purchases-load", state)

	notifier.ToPool(n)

	if flawError != nil {
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = nil
	state.Response.Body = state.NetworkAccountPurchasesLoad

	state.Response.Accepted()
}
