package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func renderTable(holdings []entities.Holding) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	table := table.New(
		"ID",
		"Name",
		"Purchase Spot Price",
		"Total Units",
		"Unit Weight",
		"Type",
	)

	table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, holding := range holdings {
		table.AddRow(
			holding.ID,
			holding.Name,
			holding.PurchaseSpotPrice,
			holding.TotalUnits,
			holding.UnitWeight,
			holding.Type,
		)
	}

	table.Print()
}

func handleListHoldings(cmd *cobra.Command, args []string) {
	db := db.GetConnection()

	var holdings []entities.Holding

	found := db.Find(&holdings)
	if found.RowsAffected < 1 {
		fmt.Print("no holdings are present, please add some. \n")
		return
	}

	renderTable(holdings)
}

var listHoldingsCmd = &cobra.Command{
	Use:   "listHoldings",
	Short: "Lists all holdings stored in the system",
	Long:  `This command lists all preciou metal holdings stored in the database.`,
	Run:   handleListHoldings,
}

func init() {
	rootCmd.AddCommand(listHoldingsCmd)
}
