package finance

import (
	"reflect"
	"testing"
	"time"
)

func TestAccount_sum(t *testing.T) {
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
			a := &Account{
				Transactions: tt.transactions,
			}
			if got := a.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTransactionAmount(amount int, date string) Transaction {
	timestamp, _ := time.Parse("2006-01-02", date)
	return Transaction{Amount: NewCurrency(amount), Date: timestamp}
}

func TestAccount_DateRange(t *testing.T) {
	account := Account{
		Transactions: Transactions{
			createTransactionAmount(500, "2017-10-01"),
			createTransactionAmount(500, "2017-10-05"),
			createTransactionAmount(500, "2017-10-17"),
			createTransactionAmount(500, "2017-10-22"),
			createTransactionAmount(500, "2017-11-04"),
		},
	}

	start, _ := time.Parse("2006-01-02", "2017-10-01")
	end, _ := time.Parse("2006-01-02", "2017-10-31")
	a := account.DateRange(start, end)

	if len(a.Transactions) != 4 {
		t.Errorf("Incorrect number of transactions identified %d %d", len(a.Transactions), 4)
	}
}

func TestAccount_StartingBalance(t *testing.T) {
	tests := []struct {
		name         string
		transactions Transactions
		want         Currency
	}{
		{
			"Starting balance is 0",
			Transactions{
				Transaction{Amount: NewCurrency(500), Balance: NewCurrency(500)},
				Transaction{Amount: NewCurrency(500), Balance: NewCurrency(1000)},
			},
			NewCurrency(0),
		},
		{
			"Starting balance is 250",
			Transactions{
				Transaction{Amount: NewCurrency(250), Balance: NewCurrency(500)},
				Transaction{Amount: NewCurrency(250), Balance: NewCurrency(750)},
				Transaction{Amount: NewCurrency(250), Balance: NewCurrency(1000)},
			},
			NewCurrency(250),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Account{
				Transactions: tt.transactions,
			}
			if got := l.StartingBalance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ledger.StartingBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
