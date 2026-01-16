package money

import (
	"errors"
	"testing"
)

func TestDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"2 decimal digts": {
			decimal:  "1.52",
			expected: Decimal{subunits: 152, precision: 2},
			err:      nil,
		},
		"not a number": {
			decimal:  "NaN",
			expected: Decimal{},
			err:      ErrInvalidDecimal,
		},
		"too large": {
			decimal:  "1234567890123",
			expected: Decimal{},
			err:      ErrTooLarge,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
