package rules

import (
	"money/finance"
	"testing"
)

func TestPersonalRules_New(t *testing.T) {
	p := New("testdata/personal-rules.json")

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

func TestRules_Rules_Apply(t *testing.T) {
	rules := Rules{
		Transactions: []TransactionRule{
			TransactionRule{
				FindReplace: "BP",
			},
		},
	}

	transactions := rules.Apply(finance.Transactions{
		finance.Transaction{Description: "BP Gas Station"},
	})

	transaction := transactions[0]
	actual := transaction.GetDescription()
	expects := "BP"

	if transaction.GetDescription() != expects {
		t.Errorf("Rules failed to apply description to transaction %s %s", expects, actual)
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

func TestRules_TransactionRule_Income(t *testing.T) {
	transaction := finance.Transaction{
		Description: "My Work",
		Amount:      finance.NewCurrency(150000),
	}

	transactionRule := TransactionRule{
		Contains: "Work",
		Income: Paycheck{
			Description: "Work Check",
			Frequency:   2,
			Amount:      finance.NewCurrency(150000),
		},
	}

	transaction, _ = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Work Check" {
		t.Errorf("Transaction description incorrect %s", transaction.GetDescription())
	}

	if transaction.Income != true {
		t.Errorf("Income wasn't identified")
	}
}

func TestRules_TransactionRule_Apply_Bill(t *testing.T) {
	transaction := finance.Transaction{
		Description: "Check #1311",
		Amount:      finance.NewCurrency(150000),
	}

	transactionRule := TransactionRule{
		Some:   []string{"check", "cashed"},
		Amount: finance.NewCurrency(150000),
		Bill: BillRule{
			Description: "Townhouse Rent",
			Day:         1,
			Amount:      finance.NewCurrency(150000),
		},
	}

	transaction, _ = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Townhouse Rent" {
		t.Errorf("Transaction description Incorrect")
	}

	if transaction.Bill != true {
		t.Errorf("Bill wasn't identified")
	}
}

func TestRules_TransactionRule_Bills(t *testing.T) {
	rules := Rules{
		Transactions: []TransactionRule{
			TransactionRule{
				Some:   []string{"check", "cashed"},
				Amount: finance.NewCurrency(150000),
				Bill: BillRule{
					Description: "Townhouse Rent",
					Day:         1,
					Amount:      finance.NewCurrency(150000),
				},
			},
			TransactionRule{
				Some:   []string{"check", "cashed"},
				Amount: finance.NewCurrency(150000),
				Bill: BillRule{
					Description: "Townhouse Rent",
					Day:         1,
					Amount:      finance.NewCurrency(150000),
				},
			},
			TransactionRule{},
			TransactionRule{},
		},
	}

	transactions := finance.Transactions{}

	bills := rules.Bills(transactions)

	if len(bills.Rules) != 2 {
		t.Error("Incorrect number of bills returned")
	}
}
