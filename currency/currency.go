package currency

import (
	"strconv"
)

// Currency is the amount of dollars represented in cents.
type Currency struct {
	Amount int
}

// New returns a currency object. It is provided the amount in cents.
func New(Amount int) Currency {
	return Currency{Amount: Amount}
}

// NewFromDollars returns a currency object. It is provided the amount in dollars.
func NewFromDollars(dollars string) Currency {
	currency := Currency{}
	currency.FromDollars(dollars)

	return currency
}

// FormatToDollars acts on the Current Type. This method outputs the currency
// amount in dollars. This method prepends the amount with the dollar sign.
func (c *Currency) FormatToDollars() string {
	return string('$') + strconv.FormatFloat(float64(c.Amount)/100, 'f', 2, 64)
}

// FromDollars accepts a string and parses it into cents.
// The expected format in dollars is #.##
func (c *Currency) FromDollars(dollars string) {
	amount, _ := strconv.ParseFloat(dollars, 32)
	amount = amount * 100.00

	// Converting the string to a float causes precision errors.
	// I correct this by rounding the amount to the nearest whole
	// number. This problem taught me that Golang doesn't include
	// a math round function.
	if amount > 0 {
		amount += 0.5
	} else {
		amount -= 0.5
	}

	c.Amount = int(amount)
}