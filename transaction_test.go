package main

import "testing"

func TestCurrency_CentsToDollars(t *testing.T) {
	type fields struct {
		Amount int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Convert 1 penny to $0.00", fields{1}, "$0.01"},
		{"Convert 120 pennies to $1.20", fields{120}, "$1.20"},
		{"Convert 1005 pennies to $10.05", fields{1005}, "$10.05"},
		{"Convert 6050400 pennies to $60504.00", fields{6050400}, "$60504.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currency := &Currency{
				Amount: tt.fields.Amount,
			}
			if got := currency.FormatToDollars(); got != tt.want {
				t.Errorf("Currency.CentsToDollars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func transactionsWithIDs(ids ...int) Transactions {
	var transactions Transactions

	for _, id := range ids {
		transactions = append(transactions, Transaction{UniqueID: id})
	}

	return transactions
}

func TestTransactions_Sort(t *testing.T) {
	type fields struct {
		Transactions
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			"Correct order of simple list",
			fields{transactionsWithIDs(2, 1, 4, 3)},
			[]int{1, 2, 3, 4},
		},
		{
			"Correct order of offset list",
			fields{transactionsWithIDs(102, 131, 422, 191)},
			[]int{102, 131, 191, 422},
		},
		{
			"Correct order with negative numbers",
			fields{transactionsWithIDs(5, 10, 2, -10)},
			[]int{-10, 2, 5, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactions := tt.fields.Transactions

			transactions.Sort()

			for i, transaction := range transactions {
				if tt.want[i] != transaction.UniqueID {

					t.Errorf("Index of [%d] = %v, want %v", i, transaction.UniqueID, tt.want[i])
				}
			}
		})
	}
}
