package bigdecimal

import (
	"strconv"
	"strings"

	"github.com/halorium/networks-example/flaw"
	"github.com/shopspring/decimal"
)

type BigDecimal struct {
	decimal.Decimal
}

func New() (*BigDecimal, *flaw.Error) {
	amount := decimal.Zero

	return &BigDecimal{amount}, nil
}

func NewFromString(val string) (*BigDecimal, *flaw.Error) {
	if strings.HasPrefix(val, ".") || val == "" {
		val = "0" + val
	}

	amount, err := decimal.NewFromString(val)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot NewFromString")
	}

	return &BigDecimal{amount}, nil
}

func NewFromFloat64(val float64) (*BigDecimal, *flaw.Error) {
	amount := decimal.NewFromFloat(val)

	return &BigDecimal{amount}, nil
}

func NewFromInt(val int) (*BigDecimal, *flaw.Error) {
	s := strconv.Itoa(val)

	return NewFromString(s)
}
