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

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.Category != "Store" && matched {
		t.Error("Category failed to assign")
	}
}

func TestRules_TransactionRule_Apply_Replace(t *testing.T) {
	transaction := finance.Transaction{Description: "CHIPOTLE 1408 ELK GROVE VI IL"}

	transactionRule := TransactionRule{
		Contains: "Chipotle",
		Replace:  "Chipotle",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Chipotle" && matched {
		t.Errorf("New description failed to replace")
	}
}

func TestRules_TransactionRule_Apply_Need(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		Contains: "BP Gas Station",
		Need:     true,
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.Need != true && matched {
		t.Error("Need wasn't assigned true")
	}
}

func TestRules_TransactionRule_Apply_FindReplace(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		FindReplace: "BP",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "BP" && matched {
		t.Error("FindReplace didnt't find or assign")
	}
}

func TestRules_TransactionRule_Apply_Remove(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		Remove: "Gas Station",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "BP" && matched {
		t.Errorf("Remove didn't work %s", transaction.GetDescription())
	}
}

func TestRules_TransactionRule_Apply_Some(t *testing.T) {
	transaction := finance.Transaction{Description: "Check #1311"}

	transactionRule := TransactionRule{
		Some:    []string{"check", "cashed"},
		Replace: "Townhouse Rent",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Townhouse Rent" && matched {
		t.Errorf("Some didn't work %s", transaction.GetDescription())
	}
}

func TestRules_TransactionRule_Apply_SomeAmount(t *testing.T) {
	transaction := finance.Transaction{
		Description: "Check #1311",
		Amount:      finance.NewCurrency(130),
	}

	transactionRule := TransactionRule{
		Some:    []string{"check", "cashed"},
		Replace: "Townhouse Rent",
		Amount:  finance.NewCurrency(150000),
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() == "Townhouse Rent" && matched {
		t.Errorf("Description should not have applied")
	}
}

func TestRules_TransactionRule_Apply_Bill(t *testing.T) {
	transaction := finance.Transaction{
		Description: "Check #1311",
		Amount:      finance.NewCurrency(150000),
	}

	transactionRule := TransactionRule{
		Some:    []string{"check", "cashed"},
		Replace: "Townhouse Rent",
		Amount:  finance.NewCurrency(150000),
		Bill: BillRule{
			Day:    1,
			Amount: finance.NewCurrency(150000),
		},
	}

	transaction, _ = transactionRule.Apply(transaction)

	if transaction.Bill != true {
		t.Errorf("Bill wasn't identified")
	}
}
