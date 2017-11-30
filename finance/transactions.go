package finance

import (
	"sort"
	"time"
)

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

// FilterByCategory reduces the list of transactions by category. Any number
// of categories can be provided.
func (t Transactions) FilterByCategory(categories ...string) Transactions {
	var output Transactions

	for _, transaction := range t {
		for _, category := range categories {
			if transaction.GetCategory() == category {
				output = append(output, transaction)
			}
		}
	}

	return output
}

// TotalExpenses is the sum of all negative amounts
func (t Transactions) TotalExpenses() Currency {
	total := 0
	for _, transaction := range t {
		if transaction.Amount.Amount < 0 {
			total += transaction.Amount.Amount
		}
	}

	return NewCurrency(total)
}

// DateRange returns Transactions between start and end time
func (t Transactions) DateRange(start time.Time, end time.Time) Transactions {
	var found Transactions

	for _, transaction := range t {
		if (transaction.Date.After(start) && transaction.Date.Before(end)) || (start == transaction.Date || end == transaction.Date) {
			found = append(found, transaction)
		}
	}

	return found
}

// GetAllByDescription returns the all transactions with a matching description
func (t Transactions) GetAllByDescription(description string) Transactions {
	foundTransactions := Transactions{}

	for _, transaction := range t {
		if transaction.GetDescription() == description {
			foundTransactions = append(foundTransactions, transaction)
		}
	}

	return foundTransactions
}

// GetByDescription returns the first transaction with a matching description
func (t Transactions) GetByDescription(description string) (Transaction, bool) {
	foundTransactions := t.GetAllByDescription(description)

	if len(foundTransactions) > 0 {
		return foundTransactions[0], true
	}

	return Transaction{}, false
}
