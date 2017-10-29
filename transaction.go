package main

import "time"

// Transaction is represents a US currency transaction.
// It is intented that all transactions reference a bank account.
// The amount is measured in Penny's.
// Currency is the amount of dollars represented in pennies.
type Currency struct {
	Amount int
}
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
