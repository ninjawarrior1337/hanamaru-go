package info

import (
	"fmt"
	"hanamaru/hanamaru"
	"hanamaru/util"
	"strconv"
)

var ExchangeRates = &hanamaru.Command{
	Name:               "xr",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Provides conversion from one currency to another: <val> <fromcurr> <tocurr (Default USD)>",
	Exec: func(ctx *hanamaru.Context) error {
		currValStr, err := ctx.GetArgIndex(0)
		if err != nil {
			return nil
		}
		currValF, err := strconv.ParseFloat(currValStr, 64)
		if err != nil {
			return nil
		}
		fromCurr, err := ctx.GetArgIndex(1)
		if err != nil {
			return nil
		}
		toCurr := ctx.GetArgIndexDefault(2, "USD")
		newVal, err := util.ConvertCurrency(currValF, fromCurr, toCurr)
		if err != nil {
			return err
		}
		ctx.Reply(fmt.Sprintf("%.2f %v -> %.2f %v", currValF, fromCurr, newVal, toCurr))
		return nil
	},
}
