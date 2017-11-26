package rules

import (
	"github.com/tomkrush/money/finance"
	"sort"
)

// Category is a list item
type Category struct {
	Category string
	Amount   finance.Currency
	Need     bool
}

// CategoryList is a list of categories
type CategoryList []Category

func (list CategoryList) get(category string) (*Category, int) {
	for i, c := range list {
		if category == c.Category {
			return &c, i
		}
	}

	return &Category{}, -1
}

// Len returns the length of a category list
func (list CategoryList) Len() int { return len(list) }

// Swap swaps to category items
func (list CategoryList) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

// Less compares category amounts to order them from largest to smallest
func (list CategoryList) Less(i, j int) bool {
	return list[i].Amount.Amount > list[j].Amount.Amount
}

func GetCategories(transactions finance.Transactions) CategoryList {
	categories := CategoryList{}

	for _, t := range transactions {
		category, i := categories.get(t.GetCategory())

		if i >= 0 {
			categories[i].Amount = categories[i].Amount.Add(t.Amount)
		} else {
			category.Category = t.GetCategory()
			category.Need = t.Need
			category.Amount = finance.NewCurrency(t.Amount.Amount)
			categories = append(categories, *category)
		}
	}

	sort.Sort(categories)

	return categories
}
