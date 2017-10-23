package main

import "testing"

func TestImporter(t *testing.T) {
	records, _ := importTransactions("testdata/transactions.csv")

	if len(records) == 0 {
		t.Errorf("Expected to import %v, but only got %v", 8, len(records))
	}
}

func TestImporterNoPath(t *testing.T) {
	_, err := importTransactions("")

	if err.Error() != "open : no such file or directory" {
		t.Error("Expected error because no path provided. %s", err.Error())
	}
}

func TestImporterIncorrectFields(t *testing.T) {
	_, err := importTransactions("testdata/transactions_error.csv")

	if err.Error() != "line 2, column 0: wrong number of fields in line" {
		t.Error("Expected error because no row field count different than header field count. %s", err.Error())
	}
}
