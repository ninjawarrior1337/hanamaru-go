package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/currency"
	"net/http"
)

var ErrIncorrectCurrency = errors.New("one of the currency symbols is incorrect")

func GetExchangeRateInfo(base currency.Unit, symbol currency.Unit) (*ExchangeRateInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s&symbols=%s", base, symbol))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var exi ExchangeRateInfo
	err = json.NewDecoder(resp.Body).Decode(&exi)
	if err != nil {
		return nil, err
	}
	if exi.Error != "" {
		return nil, errors.New(exi.Error)
	}
	return &exi, nil
}

func ConvertCurrency(fromAmt float64, from string, to string) (float64, error) {
	fromUnit, err := currency.ParseISO(from)
	toUnit, err := currency.ParseISO(to)
	if err != nil {
		return 0, err
	}

	if len(from) != 3 || len(to) != 3 {
		return 0, ErrIncorrectCurrency
	}

	exi, err := GetExchangeRateInfo(fromUnit, toUnit)
	if err != nil {
		return 0, err
	}

	return fromAmt * exi.Rates[toUnit], nil
}
