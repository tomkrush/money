package main

// Ledger is a collection of transactions
// It is expected that an operation can be performed against the entire
// set of transactions.
type Ledger struct {
	transactions Transactions
}

// Sum acts on the Ledger type to aggregate the sum of all
// transaction amounts.
func (l *Ledger) Sum() Currency {
	total := 0

	for _, t := range l.transactions {
		total += t.Amount.Amount
	}

	return Currency{total}
}
