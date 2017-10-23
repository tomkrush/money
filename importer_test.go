package main

import "testing"

func TestImporter(t *testing.T) {
	records := importTransactions("testdata/transactions.csv")

	if len(records) == 0 {
		t.Errorf("Expected to import %v, but only got %v", 8, len(records))
	}
}

func TestImporterNoPath(t *testing.T) {
	records := importTransactions("")

	if len(records) != 0 {
		t.Error("Expected to import nothing")
	}
}

func TestImporterIncorrectFields(t *testing.T) {
	importTransactions("testdata/transactions_error.csv")
}
