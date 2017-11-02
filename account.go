package main

type Account struct {
	Transactions  Transactions
	AccountNumber string
}

func (a *Account) Sum() Currency {
	total := 0

	for _, t := range a.Transactions {
		total += t.Amount.Amount
	}

	return Currency{total}
}

func (a *Account) StartingBalance() Currency {
	t := a.Transactions[0]

	return Currency{t.Balance.Amount - t.Amount.Amount}
}
