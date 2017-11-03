package main

import (
	"reflect"
	"testing"
	"time"
)

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

func TestTransactions_SplitIntoAccounts(t *testing.T) {
	transactions := Transactions{
		Transaction{Amount: Currency{5}, AccountNumber: "1"},
		Transaction{Amount: Currency{7}, AccountNumber: "1"},
		Transaction{Amount: Currency{10}, AccountNumber: "2"},
		Transaction{Amount: Currency{75}, AccountNumber: "2"},
		Transaction{Amount: Currency{100}, AccountNumber: "3"},
	}

	accounts := transactions.SplitIntoAccounts()

	if len(accounts) != 3 {
		t.Errorf("Transactions.SplitIntoAccounts() = %v, want %v", len(accounts), 3)
	}

	account, ok := accounts["1"]

	if ok == false {
		t.Error("Account 1 should exist")
	}

	if len(account.Transactions) != 2 {
		t.Errorf("Account Transitions Length = %v, want %v", len(account.Transactions), 2)
	}
}

func TestTransactions_Sum(t *testing.T) {
	tests := []struct {
		name         string
		transactions Transactions
		want         Currency
	}{
		{
			"Sum of transactions should be 1000",
			Transactions{
				Transaction{Amount: Currency{500}, Balance: Currency{500}},
				Transaction{Amount: Currency{500}},
			},
			Currency{1000},
		},
		{
			"Sum of transactions should be 750",
			Transactions{
				Transaction{Amount: Currency{250}, Balance: Currency{250}},
				Transaction{Amount: Currency{250}},
				Transaction{Amount: Currency{250}},
			},
			Currency{750},
		},
		{
			"Sum of transactions should be 90",
			Transactions{
				Transaction{Amount: Currency{5}, Balance: Currency{5}},
				Transaction{Amount: Currency{10}},
				Transaction{Amount: Currency{75}},
			},
			Currency{90},
		},
		{
			"Sum of transactions should be 1",
			Transactions{
				Transaction{Amount: Currency{1}, Balance: Currency{1}},
			},
			Currency{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transactions.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transactions.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
