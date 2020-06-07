package util

import (
	"bytes"
	"encoding/json"
	"golang.org/x/text/currency"
	"testing"
)

func TestCurrencyStruct(t *testing.T) {
	testInput := bytes.NewBuffer([]byte(`{"rates":{"USD":0.0091540761},"base":"JPY","date":"2020-06-05"}`))
	var xrr ExchangeRateInfo
	json.Unmarshal(testInput.Bytes(), &xrr)
	if xrr.Base != currency.JPY {
		t.Error("base doesnt equal JPY when its supposed to")
		return
	}
	if xrr.Rates[currency.USD] != 0.0091540761 {
		t.Error("USD value incorrect")
		return
	}
}

//TODO: Mock the currency API
