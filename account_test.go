package main

import (
	"reflect"
	"testing"
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
				Transaction{Amount: Currency{500}},
				Transaction{Amount: Currency{500}},
			},
			Currency{1000},
		},
		{
			"Sum of transactions should be 750",
			Transactions{
				Transaction{Amount: Currency{250}},
				Transaction{Amount: Currency{250}},
				Transaction{Amount: Currency{250}},
			},
			Currency{750},
		},
		{
			"Sum of transactions should be 90",
			Transactions{
				Transaction{Amount: Currency{5}},
				Transaction{Amount: Currency{10}},
				Transaction{Amount: Currency{75}},
			},
			Currency{90},
		},
		{
			"Sum of transactions should be 1",
			Transactions{
				Transaction{Amount: Currency{1}},
			},
			Currency{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Account{
				Transactions: tt.transactions,
			}
			if got := l.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ledger.Sum() = %v, want %v", got, tt.want)
			}
		})
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
				Transaction{Amount: Currency{500}, Balance: Currency{500}},
				Transaction{Amount: Currency{500}, Balance: Currency{1000}},
			},
			Currency{0},
		},
		{
			"Starting balance is 250",
			Transactions{
				Transaction{Amount: Currency{250}, Balance: Currency{500}},
				Transaction{Amount: Currency{250}, Balance: Currency{750}},
				Transaction{Amount: Currency{250}, Balance: Currency{1000}},
			},
			Currency{250},
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
