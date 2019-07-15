package currency

import (
	bigdecimal "github.com/halorium/networks-example/big-decimal"
	"github.com/halorium/networks-example/flaw"
)

func (a *Amount) DivideByInt(value int) (*Amount, *flaw.Error) {
	fromValue, flawError := bigdecimal.NewFromInt(value)

	if flawError != nil {
		return nil, flawError.Wrap("cannot DivideByInt")
	}

	newDecimal := a.Value.Div(fromValue.Abs())

	newBigDecimal, flawError := bigdecimal.NewFromString(newDecimal.String())

	if flawError != nil {
		return nil, flawError.Wrap("cannot DivideByInt")
	}

	return &Amount{newBigDecimal, a.Code}, nil
}
