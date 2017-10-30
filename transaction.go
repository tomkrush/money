package main

import (
	"strconv"
	"time"
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

// Transaction is the withdrawl and deposit of a currency from a
// bank account on a specific date.
type Transaction struct {
	BankID        string
	AccountNumber string
	AccountType   string
	Balance       Currency
	StartDate     time.Time
	EndDate       time.Time
	Type          string
	Date          time.Time
	Amount        Currency
	UniqueID      int
	Description   string
}

// Transactions is a collection of type Transaction
type Transactions []Transaction
