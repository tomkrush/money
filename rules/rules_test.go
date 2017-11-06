package rules

import (
	"money/finance"
	"testing"
)

func TestPersonalRules_New(t *testing.T) {
	p := New("testdata/personal-rules.json")

	if p.MonthlyIncome != 3000 {
		t.Error("Monthly Income incorrect")
	}

	if len(p.Categories) == 0 {
		t.Error("No categories imported")
	}

	if len(p.Transactions) == 0 {
		t.Error("No transactions imported")
	}
}

func TestRules_TransactionRule_Apply_Category(t *testing.T) {
	transaction := finance.Transaction{Description: "Target"}

	transactionRule := TransactionRule{
		Contains: "Target",
		Category: "Store",
	}

	transaction = transactionRule.Apply(transaction)

	if transaction.Category != "Store" {
		t.Error("Category failed to assign")
	}
}

func TestRules_TransactionRule_Apply_Replace(t *testing.T) {
	transaction := finance.Transaction{Description: "Target – Somewhere"}

	transactionRule := TransactionRule{
		Contains: "Target",
		Replace:  "Target",
	}

	transaction = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Target" {
		t.Error("New description failed to replace")
	}
}

func TestRules_TransactionRule_Apply_Need(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		Contains: "BP Gas Station",
		Need:     true,
	}

	transaction = transactionRule.Apply(transaction)

	if transaction.Need != true {
		t.Error("Need wasn't assigned true")
	}
}

func TestRules_TransactionRule_Apply_FindReplace(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		FindReplace: "BP",
	}

	transaction = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "BP" {
		t.Error("FindReplace didnt't find or assign")
	}
}

func TestRules_TransactionRule_Apply_Remove(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		Remove: "Gas Station",
	}

	transaction = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "BP" {
		t.Errorf("Remove didn't work %s", transaction.GetDescription())
	}
}

func TestRules_TransactionRule_Apply_Some(t *testing.T) {
	transaction := finance.Transaction{Description: "Check #1311"}

	transactionRule := TransactionRule{
		Some:    []string{"check", "cashed"},
		Replace: "Townhouse Rent",
	}

	transaction = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Townhouse Rent" {
		t.Errorf("Some didn't work %s", transaction.GetDescription())
	}
}
