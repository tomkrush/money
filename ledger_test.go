package main

import (
	"reflect"
	"testing"
)

func transactionsWithAmounts(amounts ...int) []Transaction {
	var transactions []Transaction

	for _, a := range amounts {
		transactions = append(transactions, Transaction{Amount: Currency{a}})
	}

	return transactions
}

func TestLedger_sum(t *testing.T) {
	type fields struct {
		transactions []Transaction
	}
	tests := []struct {
		name   string
		fields fields
		want   Currency
	}{
		{
			"Sum of transactions should be 1000",
			fields{transactionsWithAmounts(500, 500)},
			Currency{1000},
		},
		{
			"Sum of transactions should be 750",
			fields{transactionsWithAmounts(250, 250, 250)},
			Currency{750},
		},
		{
			"Sum of transactions should be 90",
			fields{transactionsWithAmounts(5, 10, 75)},
			Currency{90},
		},
		{
			"Sum of transactions should be 1",
			fields{transactionsWithAmounts(1)},
			Currency{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Ledger{
				transactions: tt.fields.transactions,
			}
			if got := l.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ledger.sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
