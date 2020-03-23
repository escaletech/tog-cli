package table

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func Render(header []string, rows [][]string) {
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(false)
	table.SetAutoWrapText(false)

	table.SetHeader(header)
	headerColors := make([]tablewriter.Colors, len(header))
	for i := range header {
		headerColors[i] = tablewriter.Colors{tablewriter.FgHiBlackColor}
	}
	table.SetHeaderColor(headerColors...)

	table.AppendBulk(rows)
	table.Render()
	fmt.Println()
}
