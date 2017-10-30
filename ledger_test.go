package main

import (
	"reflect"
	"testing"
)

func transactionsWithAmounts(amounts ...int) Transactions {
	var transactions Transactions

	for _, a := range amounts {
		transactions = append(transactions, Transaction{Amount: Currency{a}})
	}

	return transactions
}

func TestLedger_sum(t *testing.T) {
	type fields struct {
		transactions Transactions
	}
	tests := []struct {
		name         string
		transactions Transactions
		want         Currency
	}{
		{
			"Sum of transactions should be 1000",
			transactionsWithAmounts(500, 500),
			Currency{1000},
		},
		{
			"Sum of transactions should be 750",
			transactionsWithAmounts(250, 250, 250),
			Currency{750},
		},
		{
			"Sum of transactions should be 90",
			transactionsWithAmounts(5, 10, 75),
			Currency{90},
		},
		{
			"Sum of transactions should be 1",
			transactionsWithAmounts(1),
			Currency{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Ledger{
				transactions: tt.transactions,
			}
			if got := l.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ledger.sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
