package finance

// Account can only contain transactions with the same account id.
//
// I should probably have a function that validates that all the account ids
// match. Right now this behavior is assumed.
type Account struct {
	Transactions  Transactions
	AccountNumber string
}

// Sum on an Account takes the sum of the transactions plus the starting balance
func (a Account) Sum() Currency {
	startingBalance := a.StartingBalance().Amount
	total := a.Transactions.Sum().Amount

	return NewCurrency(startingBalance + total)
}

// StartingBalance is the balance of the first transaction minus the
// transaction amount
func (a *Account) StartingBalance() Currency {
	t := a.Transactions[0]

	return NewCurrency(t.Balance.Amount - t.Amount.Amount)
}
