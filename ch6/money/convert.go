package money

// ExchangeRate represents a rate to convert from one currency to another
type ExchangeRate Decimal

func multiply(d Decimal, r ExchangeRate) Decimal {
	dec := Decimal{
		subunits:  d.subunits * r.subunits,
		precision: d.precision + r.precision,
	}
	// Let's clean the representation a bit. Remove trailing zeros.
	dec.simplify()
	return dec
}

func Convert(amount Amount, to Currency) (Amount, error) {
	// Convert to the target currency applying the fetched change rate.
	convertedValue := applyExchangeRate(amount, to, ExchangeRate(Decimal{subunits: 2, precision: 0}))
	// Validate the converted amount is in the handled bounded range.
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}
	return convertedValue, nil
}
