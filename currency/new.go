package currency

import (
	"github.com/halorium/networks-example/big-decimal"
	"github.com/halorium/networks-example/flaw"
)

type Amount struct {
	Value *bigdecimal.BigDecimal
	Code  string
}

func NewAmount(code string) (*Amount, *flaw.Error) {
	bigDecimal, err := bigdecimal.New()

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot NewAmount")
	}

	flawError := validate(code)

	if flawError != nil {
		return nil, flawError.Wrap("currency-new-amount")
	}

	return &Amount{bigDecimal, code}, nil
}

func NewAmountFromString(val string, code string) (*Amount, *flaw.Error) {
	bigDecimal, err := bigdecimal.NewFromString(val)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot NewAmountFromString")
	}

	flawError := validate(code)

	if flawError != nil {
		return nil, flawError.Wrap("currency-new-amount-from-string")
	}

	return &Amount{bigDecimal, code}, nil
}

func NewAmountFromFloat(val float64, code string) (*Amount, *flaw.Error) {
	bigDecimal, err := bigdecimal.NewFromFloat64(val)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot NewAmountFromFloat")
	}

	flawError := validate(code)

	if flawError != nil {
		return nil, flawError.Wrap("currency-new-amount-from-float")
	}

	return &Amount{bigDecimal, code}, nil
}

func NewAmountFromInt(val int, code string) (*Amount, *flaw.Error) {
	bigDecimal, err := bigdecimal.NewFromInt(val)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot NewAmountFromInt")
	}

	flawError := validate(code)

	if flawError != nil {
		return nil, flawError.Wrap("currency-new-amount-from-int")
	}

	return &Amount{bigDecimal, code}, nil
}
