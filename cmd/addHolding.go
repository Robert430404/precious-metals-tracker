package cmd

import (
	"fmt"

	"github.com/robert430404/precious-metals-tracker/entities"
	"github.com/spf13/cobra"
)

func HandleAddHolding(cmd *cobra.Command, args []string) {
	holding := entities.Holding{}

	holding.Hydrate()

	fmt.Printf("addHolding called %q, %q, %q, %q, %q, %q\n", holding.Price, holding.Source, holding.PurchaseSpot, holding.Volume, holding.Weight, holding.Type)
}

var addHoldingCmd = &cobra.Command{
	Use:   "addHolding",
	Short: "Adds a precious metals holding.",
	Long: `This command walks you through adding a precious metals holding.

It requests the following information:

	- Purchase Price
	- Purchase Source
	- Spot Price (at time of purchase)
	- How Many Units
	- Weight Of A Single Unit (in toz)
	- Type Of Holding [Gold, Silver]

This then stores it for use inside of the tool.`,
	Run: HandleAddHolding,
}

func init() {
	rootCmd.AddCommand(addHoldingCmd)
}
