package main

import (
	"money/currency"
	"reflect"
	"testing"
)

func TestAccount_sum(t *testing.T) {
	tests := []struct {
		name         string
		transactions Transactions
		want         currency.Currency
	}{
		{
			"Sum of transactions should be 1000",
			Transactions{
				Transaction{Amount: currency.New(500), Balance: currency.New(500)},
				Transaction{Amount: currency.New(500)},
			},
			currency.New(1000),
		},
		{
			"Sum of transactions should be 750",
			Transactions{
				Transaction{Amount: currency.New(250), Balance: currency.New(250)},
				Transaction{Amount: currency.New(250)},
				Transaction{Amount: currency.New(250)},
			},
			currency.New(750),
		},
		{
			"Sum of transactions should be 90",
			Transactions{
				Transaction{Amount: currency.New(5), Balance: currency.New(5)},
				Transaction{Amount: currency.New(10)},
				Transaction{Amount: currency.New(75)},
			},
			currency.New(90),
		},
		{
			"Sum of transactions should be 1",
			Transactions{
				Transaction{Amount: currency.New(1), Balance: currency.New(1)},
			},
			currency.New(1),
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

func TestAccount_StartingBalance(t *testing.T) {
	tests := []struct {
		name         string
		transactions Transactions
		want         currency.Currency
	}{
		{
			"Starting balance is 0",
			Transactions{
				Transaction{Amount: currency.New(500), Balance: currency.New(500)},
				Transaction{Amount: currency.New(500), Balance: currency.New(1000)},
			},
			currency.New(0),
		},
		{
			"Starting balance is 250",
			Transactions{
				Transaction{Amount: currency.New(250), Balance: currency.New(500)},
				Transaction{Amount: currency.New(250), Balance: currency.New(750)},
				Transaction{Amount: currency.New(250), Balance: currency.New(1000)},
			},
			currency.New(250),
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
