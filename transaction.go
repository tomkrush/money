package main

import "time"

// Currency is the amount of dollars represented in pennies.
type Currency struct {
	Amount int
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
	UniqueID      string
	Description   string
}
