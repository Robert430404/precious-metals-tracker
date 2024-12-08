package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func renderValueTable(value string, spotPrice string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

	table := table.New(
		"Current Value",
		"Current Spot Price",
	)

	table.WithHeaderFormatter(headerFmt)
	table.AddRow(value, spotPrice)

	table.Print()
}

func handleShowValue(cmd *cobra.Command, args []string) {
	repository := pricing.GetPricingRepository()
	db := db.GetConnection()

	var holdings []entities.Holding

	found := db.Find(&holdings)
	if found.RowsAffected < 1 {
		fmt.Print("no holdings are present, please add some. \n")
		return
	}

	var totalWeight float64 = 0
	for _, holding := range holdings {
		unitWeight, err := strconv.ParseFloat(holding.UnitWeight, 16)
		if err != nil {
			continue
		}

		totalUnits, err2 := strconv.ParseFloat(holding.TotalUnits, 16)
		if err2 != nil {
			continue
		}

		totalWeight += unitWeight * totalUnits
	}

	spotPrice := repository.GetSilverSpot()
	totalValue := totalWeight * spotPrice

	renderValueTable(fmt.Sprintf("$%.2f", totalValue), fmt.Sprintf("$%.2f", spotPrice))
}

var showValueCmd = &cobra.Command{
	Use:   "showValue",
	Short: "Shows a calculation of current holding value",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: handleShowValue,
}

func init() {
	rootCmd.AddCommand(showValueCmd)
}
