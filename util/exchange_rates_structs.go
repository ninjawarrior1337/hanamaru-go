package util

import (
	"encoding/json"
	"golang.org/x/text/currency"
)

type ExchangeRateInfo struct {
	Error string `json:"error"`
	Rates map[currency.Unit]float64
	Base  currency.Unit
	Date  string `json:"date"`
}

func (e *ExchangeRateInfo) UnmarshalJSON(data []byte) error {
	type Alias ExchangeRateInfo
	aux := &struct {
		Base  string             `json:"base"`
		Rates map[string]float64 `json:"rates"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	e.Base = currency.MustParseISO(aux.Base)
	e.Rates = make(map[currency.Unit]float64)
	for k, v := range aux.Rates {
		e.Rates[currency.MustParseISO(k)] = v
	}
	return nil
}
