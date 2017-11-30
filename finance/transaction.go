package finance

import (
	"strings"
	"time"
)

// Transaction is the withdrawal and deposit of a currency from a
// bank account on a specific date.
type Transaction struct {
	BankID          string
	AccountNumber   string
	AccountType     string
	Balance         Currency
	StartDate       time.Time
	EndDate         time.Time
	Type            string
	Date            time.Time
	Amount          Currency
	UniqueID        int
	Description     string
	UserDescription string
	Category        string
	Need            bool
	Bill            bool
	Income          bool
}

// GetDescription will return the user description if available, otherwise
// the preset description will be returned instead.
func (t Transaction) GetDescription() string {
	description := t.Description

	if t.UserDescription != "" {
		description = t.UserDescription
	}

	return strings.TrimSpace(description)
}

// GetCategory returns the category if available. Otherwise, the category
// "Uncategorized" is returned.
func (t Transaction) GetCategory() string {
	category := "Uncategorized"

	if t.Category != "" {
		category = t.Category
	}

	return strings.TrimSpace(category)
}
