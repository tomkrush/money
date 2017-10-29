package main

// Ledger is a collection of transactions
// It is expected that an operation can be performed against the entire
// set of transactions.
type Ledger struct {
	transactions []Transaction
}

func (l *Ledger) sum() Currency {
	total := 0

	for _, t := range l.transactions {
		total += t.Amount.Amount
	}

	return Currency{total}
}
