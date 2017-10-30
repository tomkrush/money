package main

// Ledger is a collection of transactions
// It is expected that an operation can be performed against the entire
// set of transactions.
// type Ledger struct {
// 	transactions Transactions
// }

// Sum acts on the Ledger type to aggregate the sum of all
// transaction amounts.
// func (l *Ledger) Sum() Currency {
// 	total := l.StartingBalance().Amount

// 	for _, t := range l.transactions {
// 		total += t.Amount.Amount
// 	}

// 	return Currency{total}
// }

// func (l *Ledger) StartingBalance() Currency {
// 	accounts := make(map[string]Currency)

// 	for _, t := range l.transactions {
// 		_, ok := accounts[t.AccountNumber]

// 		if ok == false {
// 			accounts[t.AccountNumber] = Currency{t.Balance.Amount - t.Amount.Amount}
// 		}
// 	}

// 	startingBalance := 0

// 	for _, a := range accounts {
// 		startingBalance += a.Amount
// 	}

// 	return Currency{Amount: startingBalance}
// }
