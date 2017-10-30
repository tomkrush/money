package main

import (
	"strconv"
)

// Currency is the amount of dollars represented in pennies.
type Currency struct {
	Amount int
}

// FormatToDollars acts on the Current Type. This method outputs the currency
// amount in dollars. This method prepends the amount with the dollar sign.
func (c *Currency) FormatToDollars() string {
	return string('$') + strconv.FormatFloat(float64(c.Amount)/100, 'f', 2, 64)
}
