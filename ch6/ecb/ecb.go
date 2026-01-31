// Package ecb
package ecb

import (
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
	"time"
)

const (
	ErrCallingServer = ecbankError("error calling server")
	path             = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
)

type Client struct {
	url    string
	client http.Client
}

func newClient(timeout time.Duration) Client {
	return Client{
		client: http.Client{Timeout: timeout},
	}
}

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {

	if c.url == "" {
		c.url = path
	}

	res, err := c.client.Get(c.url)
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
