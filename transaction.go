package main

import (
	"money/currency"
	"time"
)

// Transaction is the withdrawl and deposit of a currency from a
// bank account on a specific date.
type Transaction struct {
	BankID          string
	AccountNumber   string
	AccountType     string
	Balance         currency.Currency
	StartDate       time.Time
	EndDate         time.Time
	Type            string
	Date            time.Time
	Amount          currency.Currency
	UniqueID        int
	Description     string
	UserDescription string
	Category        string
	Need            bool
}

// GetDescription will return the user description if available, otherwise
// the preset description will be returned instead.
func (t *Transaction) GetDescription() string {
	if t.UserDescription != "" {
		return t.UserDescription
	}

	return t.Description
}
