package money

import (
	"errors"
	"testing"
)

func TestParseCurrency(t *testing.T) {
	tt := map[string]struct {
		in       string
		expected Currency
	}{
		"hundredth EUR": {
			in:       "EUR",
			expected: Currency{code: "EUR", precision: 2}},
	}
	for name, tc := range tt {

		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.in)
			if err != nil {
				t.Errorf("expected error %v, got %v", ErrInvalidCurrencyCode, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestParseCurrency_UnknownCurrency(t *testing.T) {
	_, err := ParseCurrency("INVALID")

	if !errors.Is(err, ErrInvalidCurrencyCode) {
		t.Errorf("expected error %s, got %v", ErrInvalidCurrencyCode, err)
	}
}
