package main

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// https://github.com/olekukonko/tablewriter

const (
	ALIGN_DEFAULT = iota
	ALIGN_CENTER
	ALIGN_RIGHT
	ALIGN_LEFT
)

func main() {
	table10()
}

func table10() {
	data := [][]string{
		{"node1.example.com", "Ready"},
		{"node2.example.com", "Ready"},
		{"node3.example.com", "Ready"},
		{"node4.example.com", "NotReady"},
	}

	//table := tablewriter.NewWriter(os.Stdout)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)

	table.SetHeader([]string{"Name", "Status"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(ALIGN_LEFT)
	table.SetAlignment(ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	fmt.Println(tableString.String())
}
