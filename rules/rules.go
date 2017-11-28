package rules

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/tomkrush/money/finance"
)

// TransactionRule describes the filters and transformations for a transaction.
type TransactionRule struct {
	Contains    string           `json:"contains,omitempty"`
	Replace     string           `json:"replace,omitempty"`
	Category    string           `json:"category,omitempty"`
	Need        bool             `json:"need,omitempty"`
	FindReplace string           `json:"findReplace,omitempty"`
	Remove      string           `json:"remove,omitempty"`
	Some        []string         `json:"some,omitempty"`
	Amount      finance.Currency `json:"amount,omitempty"`
	Bill        BillRule         `json:"bill,omitempty"`
	Income      Paycheck         `json:"income,omitempty"`
}

// BillRule describes the transaction bills rules.
// Any transaction that is encountered will a Bill Rule will be marked "paid"
// The day and amount are used to drive the report predictions. The day and amount
// are not actuall used once a bill is paid. The day and amount will be automatically
// set to the actual transaction day and amount.
type BillRule struct {
	Description string           `json:"description,omitempty"`
	Day         int              `json:"day,omitempty"`
	Amount      finance.Currency `json:"amount,omitempty"`
}

// PlannedExpense allows for planned expenses to be figured into reports
type PlannedExpense struct {
	Description string           `json:"description,omitempty"`
	Amount      finance.Currency `json:"amount,omitempty"`
	Paid        time.Time        `json:"paid,omitempty"`
}

func (pe PlannedExpense) PaidClass() string {
	if pe.Paid.IsZero() == false {
		return "expense"
	}

	return ""
}

// PlannedExpenses is a slice of PlannedExpense
type PlannedExpenses []PlannedExpense

func (pe PlannedExpenses) InMonth(t time.Time) PlannedExpenses {
	expenses := PlannedExpenses{}

	for _, e := range pe {
		if e.Paid.IsZero() {
			expenses = append(expenses, e)
		} else if e.Paid.Month() == t.Month() && e.Paid.Year() == t.Year() {
			expenses = append(expenses, e)
		}
	}

	return expenses
}

// TotalExpenses is the total amount of planned expenses
func (pe PlannedExpenses) TotalExpenses() finance.Currency {
	total := 0

	for _, plannedExpense := range pe {
		if plannedExpense.Paid.IsZero() {
			total += plannedExpense.Amount.Amount
		}
	}

	return finance.NewCurrency(total)
}

// Rules contain the structures required to personalize the transaction data
// to the family needs.
type Rules struct {
	PlannedExpenses PlannedExpenses   `json:"plannedExpenses"`
	Transactions    []TransactionRule `json:"transactions"`
}

// New create the rules from a json file
func New(path string) Rules {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	rules := Rules{}

	err = json.Unmarshal(b, &rules)

	if err != nil {
		panic(err)
	}

	return rules
}

// Apply rules to transactions
func (r Rules) Apply(transactions finance.Transactions) finance.Transactions {
	for i, t := range transactions {
		for _, tr := range r.Transactions {
			t, matched := tr.Apply(t)

			if matched {
				transactions[i] = t
			}
		}
	}

	return transactions
}

func descriptionContains(transaction finance.Transaction, contains string) bool {
	description := strings.ToLower(transaction.Description)
	contains = strings.ToLower(contains)

	if contains != "" {
		if strings.Contains(description, contains) {
			return true
		}
	}

	return false
}

// Apply takes a TransactionRule and applies transformations on the transaction
// object.
//
// - Checks if a string exists in description. This is a matching property.
// - Replace will replace the string in the description
// - Category is assigned if a matching property is found
// - Need is assigned if a matched property is found
// - FindReplace checks if string exists, and replaces description if it does.
// - Remove removes the string from the description if a match is found
// - Some checks if a string in some exists. This is a matching property.
// - Amount checks if the amount of the transaction matches. If not, rule is void.
func (r TransactionRule) Apply(transaction finance.Transaction) (finance.Transaction, bool) {
	matched := false

	if r.Amount.Amount != 0 {
		if r.Amount.Amount != transaction.Amount.Amount {
			return transaction, false
		}
	}

	if descriptionContains(transaction, r.Contains) {
		matched = true
	}

	if len(r.Some) > 0 {
		for _, item := range r.Some {
			if descriptionContains(transaction, item) {
				matched = true
			}
		}
	}

	if descriptionContains(transaction, r.FindReplace) {
		matched = true
	}

	if descriptionContains(transaction, r.Remove) {
		matched = true
	}

	if matched {
		if r.Category != "" {
			transaction.Category = r.Category
		}

		if r.Replace != "" {
			transaction.UserDescription = r.Replace
		}

		if r.FindReplace != "" {
			transaction.UserDescription = r.FindReplace
		}

		if r.Remove != "" {
			transaction.UserDescription = strings.Replace(transaction.Description, r.Remove, "", -1)
		}

		transaction.Need = r.Need

		if r.Bill.Day != 0 {
			transaction.Bill = true
			transaction.UserDescription = r.Bill.Description
		}

		if r.Income.Frequency != 0 {
			transaction.Income = true
			transaction.UserDescription = r.Income.Description
		}
	}

	return transaction, matched
}

// Bills returns a list of all bills from the transaction rules
func (r Rules) Bills(transactions finance.Transactions) Bills {
	bills := Bills{}
	bills.Transactions = transactions

	for _, rule := range r.Transactions {
		if rule.Bill.Description != "" {
			bills.Rules = append(bills.Rules, rule)
		}
	}

	return bills
}
