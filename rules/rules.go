package rules

import (
	"encoding/json"
	"io/ioutil"
	"money/finance"
	"strings"
)

type TransactionRule struct {
	Contains    string   `json:"contains,omitempty"`
	Replace     string   `json:"replace,omitempty"`
	Category    string   `json:"category,omitempty"`
	Need        bool     `json:"need,omitempty"`
	FindReplace string   `json:"findReplace,omitempty"`
	Remove      string   `json:"remove,omitempty"`
	Some        []string `json:"some,omitempty"`
}

type CategoryRule struct {
	Find    string   `json:"find"`
	Replace string   `json:"replace"`
	Some    []string `json:"some"`
}

type Rules struct {
	MonthlyIncome int               `json:"monthlyIncome"`
	Categories    []CategoryRule    `json:"categories"`
	Transactions  []TransactionRule `json:"transactions"`
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

func (r TransactionRule) Apply(transaction finance.Transaction) finance.Transaction {
	matched := false

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
	}

	return transaction
}
