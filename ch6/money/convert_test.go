package money_test

import (
	"learngo-pockets/moneyconverter/money"
	"reflect"
	"testing"
)

type stubRate struct {
	rate string
	err  error
}

func (sr stubRate) FetchExchangeRate(_, _ money.Currency) (money.ExchangeRate, error) {
	rate, _ := money.ParseDecimal(sr.rate)
	return money.ExchangeRate(rate), sr.err
}

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		to       money.Currency
		rate     stubRate
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.96 BHD to EUR": {
			amount: mustParseAmount(t, "34.98", "BHD"),
			to:     MustParseCurrency(t, "EUR"),
			rate:   stubRate{rate: "2"},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error go %s", err.Error())
				}
				expected := mustParseAmount(t, "69.96", "EUR")
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("Exepted %v , got %v", expected, got)
				}
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to, tc.rate)
			tc.validate(t, got, err)
		})
	}
}

func MustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()
	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("Cannot parse currency %s code", code)
	}
	return currency
}

func mustParseAmount(t *testing.T, value string, code string) money.Amount {
	t.Helper()
	n, err := money.ParseDecimal(value)
	if err != nil {
		t.Fatalf("invalid number: %s", value)
	}
	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("invalid currency code: %s", code)
	}
	amount, err := money.NewAmount(n, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %v and currency code %s", n, code)
	}
	return amount
}
