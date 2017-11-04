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

// Sum acts on the Ledger type to aggregate the sum of all
// transaction amounts.
func (t Transactions) Sum() Currency {
	total := 0

	for _, item := range t {
		total += item.Amount.Amount
	}

	return Currency{total}
}
