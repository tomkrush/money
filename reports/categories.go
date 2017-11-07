package reports

import (
	"fmt"
	"money/finance"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

type Category struct {
	Category string
	Amount   int
	Need     bool
}

type CategoryList []Category

func (list CategoryList) get(category string) (*Category, int) {
	for i, c := range list {
		if category == c.Category {
			return &c, i
		}
	}

	return &Category{}, -1
}

func (p CategoryList) Len() int           { return len(p) }
func (p CategoryList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p CategoryList) Less(i, j int) bool { return p[i].Amount < p[j].Amount }

func Categories(transactions finance.Transactions) {
	fmt.Println("# Categories")
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
