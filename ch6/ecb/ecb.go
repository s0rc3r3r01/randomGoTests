// Package ecb
package ecb

import (
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
)

const (
	ErrCallingServer = ecbankError("error calling server")
	path             = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
)

type Client struct {
}

func (c Client) FetchExchagngeeRate(source, target money.Currency) (money.ExchangeRate, error) {

	res, err := http.Get(path)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %d", ErrCallingServer, err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return money.ExchangeRate{}, fmt.Errorf("%w: %d", ErrCallingServer, res.StatusCode)
	}

	rate, err := readRateFromResponse(source.GetCode(), target.GetCode(), res.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil

}
