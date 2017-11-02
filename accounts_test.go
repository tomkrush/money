package main

import "testing"

func TestAccounts_transactionstoAccounts(t *testing.T) {
	transactions := Transactions{
		Transaction{Amount: Currency{5}, AccountNumber: "1"},
		Transaction{Amount: Currency{7}, AccountNumber: "1"},
		Transaction{Amount: Currency{10}, AccountNumber: "2"},
		Transaction{Amount: Currency{75}, AccountNumber: "2"},
		Transaction{Amount: Currency{100}, AccountNumber: "3"},
	}

	accounts := transactions.SplitIntoAccounts(transactions)

	if len(accounts) != 3 {
		t.Errorf("Transactions.SplitIntoAccounts() = %v, want %v", len(accounts), 3)
	}

	account, ok := accounts["1"]

	if ok == false {
		t.Error("Account 1 should exist")
	}

	if len(account.Transactions) == 2 {
		t.Errorf("Account Transitions Length = %v, want %v", len(account.Transactions), 2)
	}
}
