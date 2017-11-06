package finance

import "sort"

// Transactions is a collection of type Transaction
type Transactions []Transaction

// SplitIntoAccounts takes a list of transactions and groups them by account.
// This is done by looking at the accountNumber on each transaction.
func (t Transactions) SplitIntoAccounts() Accounts {
	accounts := make(Accounts)

	for _, item := range t {
		accountNumber := item.AccountNumber
		account := accounts[accountNumber]
		account.AccountNumber = item.AccountNumber
		account.Transactions = append(account.Transactions, item)
		accounts[accountNumber] = account
	}

	return accounts
}

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
	total := 0

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
