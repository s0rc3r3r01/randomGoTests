package money

import "fmt"

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

func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get change rate %w", err)
	}

	// Convert to the target currency applying the fetched change rate.
	convertedValue := applyExchangeRate(amount, to, r)
	// Validate the converted amount is in the handled bounded range.
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}
	return convertedValue, nil
}
