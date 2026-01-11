package money

import (
	"errors"
	"testing"
)

func TestAmount(t *testing.T) {
	tt := map[string]struct {
		testedQuantity Decimal
		testedCurrency Currency
		expected       Amount
		err            error
	}{
		"precise good precision": {
			testedQuantity: Decimal{subunits: 15, precision: 2},
			testedCurrency: Currency{code: "EUR", precision: 2},
			expected:       Amount{quantity: Decimal{subunits: 15, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			err:            nil,
		},
		"precisionTooHigh": {
			testedQuantity: Decimal{subunits: 15, precision: 4},
			testedCurrency: Currency{code: "EUR", precision: 2},
			expected:       Amount{},
			err:            ErrTooPrecise,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.testedQuantity, tc.testedCurrency)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
