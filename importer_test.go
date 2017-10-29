package main

import (
	"testing"
)

func TestImporter(t *testing.T) {
	transactions, _ := importTransactions("testdata/transactions.csv")

	if len(transactions) == 0 {
		t.Errorf("Expected to import %v, but only got %v", 8, len(transactions))
	}
}

func TestFirstRecordData(t *testing.T) {
	transactions, _ := importTransactions("testdata/transactions.csv")
	transaction := transactions[0]
	bankID := "1234567"
	accountNumber := "555555555"
	accountType := "CHECKING"
	balance := 100007
	startDate := "2017-01-01"
	endDate := "2017-10-22"
	transactionType := "DEBIT"
	date := "2017-01-01"
	amount := -2250
	uniqueID := "1"
	description := "BP Gas Station"

	if transaction.BankID != bankID {
		t.Errorf("Expected bank id %s, but only got %s", bankID, transaction.BankID)
	}

	if transaction.AccountNumber != accountNumber {
		t.Errorf("Expected account number %s, but only got %s", accountNumber, transaction.AccountNumber)
	}

	if transaction.AccountType != accountType {
		t.Errorf("Expected account type %s, but only got %s", accountType, transaction.AccountType)
	}

	if transaction.AccountType != accountType {
		t.Errorf("Expected account type %s, but only got %s", accountType, transaction.AccountType)
	}

	if transaction.Balance != balance {
		t.Errorf("Expected balance %d, but only got %d", balance, transaction.Balance)
	}

	if transaction.StartDate != startDate {
		t.Errorf("Expected start date %s, but only got %s", startDate, transaction.StartDate)
	}

	if transaction.EndDate != endDate {
		t.Errorf("Expected end date %s, but only got %s", endDate, transaction.EndDate)
	}

	if transaction.Type != transactionType {
		t.Errorf("Expected transaction type %s, but only got %s", transactionType, transaction.Type)
	}

	if transaction.Date != date {
		t.Errorf("Expected date %s, but only got %s", date, transaction.Date)
	}

	if transaction.Amount != amount {
		t.Errorf("Expected amount %d, but only got %d", amount, transaction.Amount)
	}

	if transaction.UniqueID != uniqueID {
		t.Errorf("Expected unique id %s, but only got %s", uniqueID, transaction.UniqueID)
	}

	if transaction.Description != description {
		t.Errorf("Expected description id %s, but only got %s", description, transaction.Description)
	}
}

func TestImporterNoPath(t *testing.T) {
	_, err := importTransactions("")

	if err.Error() != "open : no such file or directory" {
		t.Errorf("Expected error because no path provided. %s", err.Error())
	}
}

func TestImporterIncorrectFields(t *testing.T) {
	_, err := importTransactions("testdata/transactions_error.csv")

	if err.Error() != "line 2, column 0: wrong number of fields in line" {
		t.Errorf("Expected error because no row field count different than header field count. %s", err.Error())
	}
}
