package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/spf13/cobra"
)

func HandleFirstRun() {
	loadedConfig := config.GetConfig()
	fmt.Printf("Derived ConfigPath: %q, %v\n", loadedConfig.ConfigPath, loadedConfig.RuntimeFlags.AddHoldingRan)
	configPath := config.GetConfig().ConfigPath

	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("there was a problem ensuring the config path: %v\n", err))
	}
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
