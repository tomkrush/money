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

func (t *Transaction) getDescription() string {
	if t.UserDescription != "" {
		return t.UserDescription
	}

	return t.Description
}
