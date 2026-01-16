package money

import (
	"strconv"
	"strings"
)

const (
	ErrInvalidDecimal = Error("unable to convert to decimal")
	ErrTooLarge       = Error("Too Large")
)
const maxDecimal = 1e12

type Decimal struct {
	subunits  int64
	precision byte
}

func ParseDecimal(value string) (Decimal, error) {

	intPart, fracPart, _ := strings.Cut(value, ".")
	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, ErrInvalidDecimal
	}
	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}
	precision := byte(len(fracPart))
	return Decimal{subunits: subunits, precision: precision}, nil
}

func (d *Decimal) simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}
