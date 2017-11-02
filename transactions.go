package main

import (
	"sort"
)

// Transactions is a collection of type Transaction
type Transactions []Transaction

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

// SplitIntoAccounts takes a list of transactions and groups them by account.
// This is done by looking at the accountNumber on each transaction.
func (t *Transactions) SplitIntoAccounts(transactions Transactions) Accounts {
	accounts := make(Accounts)

	for _, t := range transactions {
		accountNumber := t.AccountNumber
		account := accounts[accountNumber]
		account.AccountNumber = t.AccountNumber
		accounts[accountNumber] = account

		account.Transactions = append(account.Transactions, t)
	}

	return accounts
}
