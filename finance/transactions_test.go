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

func TestTransactions_SplitIntoAccounts(t *testing.T) {
	transactions := Transactions{
		Transaction{Amount: NewCurrency(5), AccountNumber: "1"},
		Transaction{Amount: NewCurrency(7), AccountNumber: "1"},
		Transaction{Amount: NewCurrency(10), AccountNumber: "2"},
		Transaction{Amount: NewCurrency(75), AccountNumber: "2"},
		Transaction{Amount: NewCurrency(100), AccountNumber: "3"},
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
