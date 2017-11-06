package finance

import "time"

// Account can only contain transactions with the same account id.
//
// I should probably have a function that validates that all the account ids
// match. Right now this behavior is assumed.
type Account struct {
	Transactions  Transactions
	AccountNumber string
}

// DateRange returns Transactions between start and end time
func (a Account) DateRange(start time.Time, end time.Time) Account {
	var found Transactions

	for _, t := range a.Transactions {
		if (t.Date.After(start) && t.Date.Before(end)) || (start == t.Date || end == t.Date) {
			found = append(found, t)
		}
	}

	a.Transactions = found

	return a
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
