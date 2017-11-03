package main

type Account struct {
	Transactions Transactions
}

func (a Account) Sum() Currency {
	startingBalance := a.StartingBalance().Amount
	total := a.Transactions.Sum().Amount

	return Currency{startingBalance + total}
}

func (a *Account) StartingBalance() Currency {
	t := a.Transactions[0]

	return Currency{t.Balance.Amount - t.Amount.Amount}
}
