package cycler

import (
	"bytes"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/http-client"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network-account-purchase-load"
	"github.com/halorium/networks-example/serializers"
	"github.com/halorium/networks-example/try"
)

func postNetworkAccountPurchasesLoad(cycle *Cycle) *flaw.Error {
	logger.Info(cycle.Tag, cycle)

	try := try.New()

	for {
		body := &networkaccountpurchaseload.Entity{
			StartTime: cycle.StartTime,
			StopTime:  cycle.StopTime,
		}

		serializedBody := serializers.JSONCompactSerializer(body)
		readerBody := bytes.NewReader(serializedBody)

		httpClient := httpclient.FromPool()

		_, flawError := flaw.FromHTTPResponse(
			httpClient.Post(cycle.URI, "application/json", readerBody),
		)

		httpclient.ToPool(httpClient)

		if flawError != nil {
			if try.Again(cycle.Tag+"-retry", flawError) {
				continue
			}

			return flawError.Wrap("cannot postNetworkAccountPurchasesLoad")
		}

		return nil
	}
}
