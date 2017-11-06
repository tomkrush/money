package finance

// Accounts is a list of values of type account
type Accounts map[string]Account

// StartingBalance provides the sum all of starting balances for each account.
func (accounts Accounts) StartingBalance() Currency {
	startingBalance := 0

	for _, a := range accounts {
		startingBalance += a.StartingBalance().Amount
	}

	return Currency{startingBalance}
}
