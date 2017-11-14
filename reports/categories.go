package reports

import (
	"fmt"
	"money/finance"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// Category is a list item
type Category struct {
	Category string
	Amount   int
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
func (list CategoryList) Less(i, j int) bool { return list[i].Amount < list[j].Amount }

// Categories displays a formatted table of categories. This report is useful
// for visualizing how transactions are grouped.
func Categories(transactions finance.Transactions) {
	fmt.Println("### Categories")
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "Amount", "Need"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	categories := CategoryList{}

	for _, t := range transactions {
		category, i := categories.get(t.GetCategory())

		if i >= 0 {
			categories[i].Amount += t.Amount.Amount
		} else {
			category.Category = t.GetCategory()
			category.Need = t.Need
			category.Amount = t.Amount.Amount
			categories = append(categories, *category)
		}
	}

	sort.Sort(categories)

	for _, category := range categories {
		amount := finance.NewCurrency(category.Amount)
		need := "No"

		if category.Need {
			need = "Yes"
		}

		table.Append([]string{
			category.Category,
			amount.FormatToDollars(),
			need,
		})
	}

	table.Render() // Send output
}
