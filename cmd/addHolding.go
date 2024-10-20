package cmd

import (
	"fmt"
	"os"

	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/robert430404/precious-metals-tracker/transformers"
	"github.com/spf13/cobra"
)

func handleFirstRun() {
	loadedConfig := config.GetConfig()
	if loadedConfig.RuntimeFlags.AddHoldingRan {
		return
	}

	configPath := config.GetConfig().ConfigPath

	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("there was a problem ensuring the config path: %v\n", err))
	}

	loadedConfig.RuntimeFlags.SetAddHoldingRan(true)
}

func handleAddHolding(cmd *cobra.Command, args []string) {
	handleFirstRun()

	holding := &models.Holding{}
	holding.Hydrate()

	db := db.GetConnection()
	transformer := transformers.HoldingTransformer{}

	db.Create(transformer.TransformModelToEntity(holding))
	fmt.Print("stored holding")
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
	Run: handleAddHolding,
}

func init() {
	rootCmd.AddCommand(addHoldingCmd)
}
