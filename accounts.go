package main

import (
	"money/currency"
)

// Accounts is a list of values of type account
type Accounts map[string]Account

// StartingBalance provides the sum all of starting balances for each account.
func (accounts Accounts) StartingBalance() currency.Currency {
	startingBalance := 0

	for _, a := range accounts {
		startingBalance += a.StartingBalance().Amount
	}

	return currency.Currency{startingBalance}
}
