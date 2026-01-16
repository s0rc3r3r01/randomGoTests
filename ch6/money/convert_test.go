package money_test

import (
	"learngo-pockets/moneyconverter/money"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.96 USD to EUR": {
			amount: mustParseAmount(t, "34.98", " USD"),
			to:     MustParseCurrency(t, "EUR"),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error go %s", err.Error())
				}
				expected := money.Amount{}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("Exepted %v , got %v", expected, got)
				}
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
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
