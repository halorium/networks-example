package main

import (
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/affiliate-window/transaction-loader"
	"github.com/halorium/networks-example/networks/affilinet/transaction-loader"
	"github.com/halorium/networks-example/networks/booking/transaction-loader"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/networks/rakuten/transaction-loader"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/state"
	"github.com/halorium/networks-example/uuid"
	"github.com/halorium/networks-example/validation"
)

func putNetworkAccountPurchase(state *state.State) {
	logger.Debug("put-network-account-purchase", state)

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

	state.NetworkAccountPurchase = networkaccountpurchase.NewEntity()

	state.NetworkAccountPurchase.Head.NetworkID = state.Request.Params.ByName("network-id")
	state.NetworkAccountPurchase.Head.NetworkAccountID = state.Request.Params.ByName("network-account-id")
	state.NetworkAccountPurchase.Head.PurchaseID = state.Request.Params.ByName("purchase-id")
	state.NetworkAccountPurchase.Head.CreatedAt = state.CreatedAt.String()

	flawError := serializers.JSONDeserializer(state.Request.Body, state.NetworkAccountPurchase.Body)

	if flawError != nil {
		state.FlawError = flawError.Wrap("cannot putNetworkAccountPurchase")
		state.Response.Head = nil
		state.Response.Body = validation.NewFrom("malformed body")
		state.Response.BadRequest()
		return
	}

	validationEntity := networkaccountpurchase.Validate(state.NetworkAccountPurchase)

	if validationEntity.HasErrors() {
		state.Response.Head = nil
		state.Response.Body = validationEntity
		state.Response.BadRequest()
		return
	}

	switch state.NetworkAccountPurchase.Body.Variant {
	case affiliatewindow.Variant:
		flawError = affiliatewindow.SetVariationVersion(state.NetworkAccountPurchase)
	case affilinet.Variant:
		flawError = affilinet.SetVariationVersion(state.NetworkAccountPurchase)
	case booking.Variant:
		flawError = booking.SetVariationVersion(state.NetworkAccountPurchase)
	case rakuten.Variant:
		flawError = rakuten.SetVariationVersion(state.NetworkAccountPurchase)
	default:
		flawError = flaw.New("invalid variant").Wrap("cannot putNetworkAccountPurchase")
	}

	if flawError != nil {
		state.FlawError = flawError
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.BadRequest()
		return
	}

	state.NetworkAccountPurchase.Head.Etag = uuid.NewHash(string(serializers.JSONCompactSerializer(state.NetworkAccountPurchase.Body)))

	n := notifier.FromPool()

	flawError = n.Notify("put-network-account-purchase", state)

	notifier.ToPool(n)

	if flawError != nil {
		state.Response.Head = nil
		state.Response.Body = nil
		state.Response.ServiceUnavailable()
		return
	}

	state.Response.Head = state.NetworkAccountPurchase.Head
	state.Response.Body = state.NetworkAccountPurchase.Body

	state.Response.Accepted()
}
