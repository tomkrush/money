package main

import "time"

// Transaction is represents a US currency transaction.
// It is intented that all transactions reference a bank account.
// The amount is measured in Penny's.
type Transaction struct {
	BankID        string
	AccountNumber string
	AccountType   string
	Balance       int
	StartDate     string
	EndDate       string
	Type          string
	Date          time.Time
	Amount        int
	UniqueID      string
	Description   string
}
