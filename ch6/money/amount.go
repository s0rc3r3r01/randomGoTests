package money

type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	ErrTooPrecise = Error("Quantity is too precise")
)

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}
	quantity.precision = currency.precision

	return Amount{quantity: quantity, currency: currency}, nil
}

func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted := multiply(a.quantity, rate)
	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}
	converted.precision = target.precision
	return Amount{
		currency: target,
		quantity: converted,
	}
}

func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}
	return nil
}

func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.code
}
