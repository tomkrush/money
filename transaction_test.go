package main

import "testing"
import "time"

func createTransaction(uniqueID int, date string) Transaction {
	timestamp, _ := time.Parse("2006-01-02", date)
	return Transaction{UniqueID: uniqueID, Date: timestamp}
}

func TestTransactions_Sort(t *testing.T) {
	tests := []struct {
		name         string
		transactions Transactions
		want         []int
	}{
		{
			"Correct order of simple list",
			Transactions{
				createTransaction(2, "2017-01-02"),
				createTransaction(1, "2017-01-01"),
				createTransaction(3, "2017-01-05"),
				createTransaction(4, "2017-01-10"),
			},
			[]int{1, 2, 3, 4},
		},
		{
			"Correct order based on date. Out of order based on priority",
			Transactions{
				createTransaction(2, "2017-01-02"),
				createTransaction(1, "2017-01-01"),
				createTransaction(422, "2017-01-01"),
				createTransaction(3, "2017-01-05"),
				createTransaction(4, "2017-01-10"),
				createTransaction(423, "2017-01-02"),
				createTransaction(424, "2017-01-11"),
			},
			[]int{1, 422, 2, 423, 3, 4, 424},
		},
		{
			"Correct order with negative numbers",
			Transactions{
				createTransaction(5, "2017-01-01"),
				createTransaction(10, "2017-01-01"),
				createTransaction(2, "2017-01-01"),
				createTransaction(-10, "2017-01-01"),
			},
			[]int{-10, 2, 5, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactions := tt.transactions

			transactions.Sort()

			for i, transaction := range transactions {
				if tt.want[i] != transaction.UniqueID {

					t.Errorf("Index of [%d] = %v, want %v", i, transaction.UniqueID, tt.want[i])
				}
			}
		})
	}
}
