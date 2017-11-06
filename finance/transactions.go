package finance

import "sort"

// Transactions is a collection of type Transaction
type Transactions []Transaction

// StartingBalance calculates the correct starting balance by iterating over
// each transaction and identify if it belongs to a different account.
// If it does, than the starting balance is added to the balance of all accounts.
func (t Transactions) StartingBalance() Currency {
	accounts := make(map[string]int)
	balance := 0

	for _, i := range t {
		_, ok := accounts[i.AccountNumber]

		if ok == false {
			accounts[i.AccountNumber] = 0
			balance += i.Balance.Amount - i.Amount.Amount
		}
	}

	return NewCurrency(balance)
}

// Sum acts on the Ledger type to aggregate the sum of all
// transaction amounts.
func (t Transactions) Sum() Currency {
	total := t.StartingBalance().Amount

	for _, item := range t {
		total += item.Amount.Amount
	}

	return NewCurrency(total)
}

// Sort exposes the StdLib Sort command without needing to import
// the package.
func (t *Transactions) Sort() {
	sort.Sort(t)
}

func (t Transactions) Len() int {
	return len(t)
}

func (t Transactions) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Transactions) Less(i, j int) bool {
	if t[i].Date.Before(t[j].Date) {
		return true
	}

	if t[i].Date.After(t[j].Date) {
		return false
	}

	return t[i].UniqueID < t[j].UniqueID
}
