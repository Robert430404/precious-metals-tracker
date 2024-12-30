package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/validations"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the application in your environment",
	Long:  `This command sets up and migrates the SQLite database for storing your holding information.`,
	Run: func(cmd *cobra.Command, args []string) {
		loadedConfig := config.GetConfig()
		sqlitePath := loadedConfig.SqlitePath

		_, err := os.Stat(sqlitePath)
		if err != nil {
			fmt.Print("creating sqlite file because it does not exist \n")
			os.Create(sqlitePath)
		}

		db := db.GetConnection()

		fmt.Print("migrating the database \n")
		db.AutoMigrate(&entities.Holding{})

		prompt := promptui.Prompt{
			Label:    "goldapi.io API Key",
			Validate: validations.ValidateString,
			Default: loadedConfig.RuntimeFlags.GoldAPIKey,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("%q failed: %v", "goldapi.io API Key", err)
			return
		}

		loadedConfig.RuntimeFlags.SetGoldAPIKey(result)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
