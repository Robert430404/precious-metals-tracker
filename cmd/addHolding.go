package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/spf13/cobra"
)

func HandleFirstRun() {
	fmt.Printf("Derived ConfigPath: %q\n", config.GetConfig().ConfigPath)
}

func HandleAddHolding(cmd *cobra.Command, args []string) {
	HandleFirstRun()

	holding := &models.Holding{}
	holding.Hydrate()

	json, err := json.Marshal(holding)
	if err != nil {
		fmt.Printf("there was a problem stringifying your holding: %v\n", err)
		return
	}

	fmt.Printf("addHolding called %q\n", json)
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
