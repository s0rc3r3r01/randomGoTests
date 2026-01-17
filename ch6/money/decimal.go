package money

import (
	"fmt"
	"math"
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

// pow10 is a quick implementation of how to raise 10 to a given power.
func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}

func (d *Decimal) String() string {
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.subunits)
	}
	centsPerUnit := pow10(d.precision)
	frac := d.subunits % centsPerUnit
	integer := d.subunits / centsPerUnit
	decimalFormat := "%d.%0" + strconv.Itoa(int(d.precision)) + "d"
	return fmt.Sprintf(decimalFormat, integer, frac)
}
