package affilinet

import (
	"fmt"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

type paginator struct {
	Page        int
	Size        int
	RecordCount int
}

func paginate(state *state.State, apiToken string) *flaw.Error {
	logger.Debug("affilinet-paginate", state)

	paginator := &paginator{
		Page: 1,
		Size: 100,
	}

	for {
		flawError := fetchPage(state, apiToken, paginator)

		if flawError != nil {
			return flawError.Wrap("cannot paginate")
		}

		res, ok := state.VariantResponse.(response)

		if !ok {
			return flaw.New(fmt.Sprintf("invalid type assertion: %T", res)).Wrap("cannot paginate")
		}

		if paginator.RecordCount == 0 {
			paginator.RecordCount = res.TotalRecords
		}

		paginator.RecordCount -= len(res.Transactions)

		if paginator.RecordCount == 0 {
			break
		}

		paginator.Page++
	}

	return nil
}
