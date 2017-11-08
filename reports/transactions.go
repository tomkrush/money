package reports

import (
	"fmt"
	"money/finance"
	"os"

	"github.com/olekukonko/tablewriter"
)

func Transactions(transactions finance.Transactions) {
	fmt.Println("# Personal Balance")
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Description", "Date", "Category", "Need", "Amount", "Total Balance"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	total := transactions.StartingBalance().Amount

	for _, t := range transactions {
		amount := t.Amount
		need := "No"

		total += amount.Amount

		if t.Need {
			need = "Yes"
		}

		totalBalance := finance.NewCurrency(total)

		table.Append([]string{
			t.GetDescription(),
			t.Date.Format("2006-01-02"),
			t.GetCategory(),
			need,
			amount.FormatToDollars(),
			totalBalance.FormatToDollars(),
		})
	}

	table.Render() // Send output

	sum := transactions.Sum()
	fmt.Println("Total: " + sum.FormatToDollars())
}
