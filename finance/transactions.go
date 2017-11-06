package finance

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

// Sum acts on the Ledger type to aggregate the sum of all
// transaction amounts.
func (t Transactions) Sum() Currency {
	total := 0

	for _, item := range t {
		total += item.Amount.Amount
	}

	return NewCurrency(total)
}
