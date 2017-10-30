package main

type Account struct {
	transactions Transactions
}

func (a *Account) Sum() Currency {
	total := 0

	for _, t := range a.transactions {
		total += t.Amount.Amount
	}

	return Currency{total}
}

func (a *Account) StartingBalance() Currency {
	t := a.transactions[0]

	return Currency{t.Balance.Amount - t.Amount.Amount}
}
