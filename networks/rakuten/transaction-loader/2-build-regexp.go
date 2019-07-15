package rakuten

import (
	"regexp"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/state"
)

func buildRegexp(state *state.State) *flaw.Error {
	logger.Debug("rakuten-build-regexp", state)

	year := state.NetworkAccountPurchasesLoad.StartTime.YYYY()             // "2006"
	longMonth := state.NetworkAccountPurchasesLoad.StartTime.LongMonth()   // "January"
	shortMonth := state.NetworkAccountPurchasesLoad.StartTime.ShortMonth() // "Jan"
	day := state.NetworkAccountPurchasesLoad.StartTime.DD()                // "02"

	username := state.NetworkAccount.Body.Variations.Username

	regexpString := "(?i)" +
		username +
		"_DailyU1_" +
		year +
		"(" +
		longMonth +
		"|" +
		shortMonth +
		")" +
		day +
		"-\\d{4}.xml"

	reg, err := regexp.Compile(regexpString)

	if err != nil {
		return flaw.From(err).Wrap("cannot buildRegexp")
	}

	state.Regexp = reg

	return openSFTPFile(state)
}
