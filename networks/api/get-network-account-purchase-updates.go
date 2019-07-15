package main

import (
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/stamps"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/validation"
)

func getNetworkAccountPurchaseUpdates(state *state.State) {
	logger.Debug("get-network-account-purchase-updates", state)

	networkAccountPurchaseEntity := networkaccountpurchase.NewEntity()
	networkAccountPurchaseEntity.Head.NetworkID = state.Request.Params.ByName("network-id")
	networkAccountPurchaseEntity.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")

	createdAtDate := state.Request.FormValue("created-at")

	_, flawError := stamps.ParseStamp(createdAtDate, stamps.YYYYhMMhDD)

	if flawError != nil {
		state.FlawError = flaw.New("created-at is invalid").Wrap("cannot getNetworkAccountPurchaseUpdates")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("created-at is invalid")
		state.Response.BadRequest()
		return
	}

	collection, flawError := entities.GetEntitiesFromPostgresByCreatedAt(networkAccountPurchaseEntity, createdAtDate)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot getNetworkAccountPurchaseUpdates")
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.NotFound()
		return
	}

	state.Response.Head = nil
	state.Response.Body = collection
	state.Response.Ok()
}
