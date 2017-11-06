package finance

import (
	"reflect"
	"testing"
	"time"
)

func createTransaction(uniqueID int, date string) Transaction {
	timestamp, _ := time.Parse("2006-01-02", date)
	return Transaction{UniqueID: uniqueID, Date: timestamp}
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
				Transaction{Amount: NewCurrency(500), Balance: NewCurrency(500)},
				Transaction{Amount: NewCurrency(500)},
			},
			NewCurrency(1000),
		},
		{
			"Sum of transactions should be 750",
			Transactions{
				Transaction{Amount: NewCurrency(250), Balance: NewCurrency(250)},
				Transaction{Amount: NewCurrency(250)},
				Transaction{Amount: NewCurrency(250)},
			},
			NewCurrency(750),
		},
		{
			"Sum of transactions should be 90",
			Transactions{
				Transaction{Amount: NewCurrency(5), Balance: NewCurrency(5)},
				Transaction{Amount: NewCurrency(10)},
				Transaction{Amount: NewCurrency(75)},
			},
			NewCurrency(90),
		},
		{
			"Sum of transactions should be 1",
			Transactions{
				Transaction{Amount: NewCurrency(1), Balance: NewCurrency(1)},
			},
			NewCurrency(1),
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

func TestTransactions_StartingBalance(t *testing.T) {
	tests := []struct {
		name         string
		transactions Transactions
		want         Currency
	}{
		{
			"Starting Balance of transactions should be 0",
			Transactions{
				Transaction{Amount: NewCurrency(500), Balance: NewCurrency(500)},
				Transaction{Amount: NewCurrency(500)},
			},
			NewCurrency(0),
		},
		{
			"Sum of transactions should be 750",
			Transactions{
				Transaction{Amount: NewCurrency(250), Balance: NewCurrency(250)},
				Transaction{Amount: NewCurrency(250)},
				Transaction{Amount: NewCurrency(250)},
			},
			NewCurrency(0),
		},
		{
			"Sum of transactions should be 90",
			Transactions{
				Transaction{Amount: NewCurrency(5), Balance: NewCurrency(10)},
				Transaction{Amount: NewCurrency(10)},
				Transaction{Amount: NewCurrency(75)},
			},
			NewCurrency(5),
		},
		{
			"Sum of transactions should be 1",
			Transactions{
				Transaction{Amount: NewCurrency(1), Balance: NewCurrency(10)},
			},
			NewCurrency(9),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transactions.StartingBalance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transactions.StartingBalance() = %v, want %v", got, tt.want)
			}
		})
	}
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
