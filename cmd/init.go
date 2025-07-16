package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db/migrations"
	"github.com/robert430404/precious-metals-tracker/validations"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [--api-key key]",
	Short: "Initializes the application in your environment",
	Long:  `This command sets up and migrates the SQLite database for storing your holding information.`,
	Run: func(cmd *cobra.Command, args []string) {
		loadedConfig, err := config.GetConfig()
		if err != nil {
			fmt.Printf("there was a problem loading your configuration: %v", err)
		}

		sqlitePath := loadedConfig.SqlitePath

		_, err = os.Stat(sqlitePath)
		if err != nil {
			fmt.Print("creating sqlite file because it does not exist \n")
			os.Create(sqlitePath)
		}

		fmt.Print("running migrations")
		migrationManager := migrations.GetMigrationsManager()
		err = migrationManager.Init()
		if err != nil {
			fmt.Printf("migrations manager failed to run migrations: %v", err)
		}

		// Check if API key was provided via CLI flag
		apiKey, err := cmd.Flags().GetString("api-key")
		if err != nil {
			fmt.Printf("there was a problem parsing the api-key flag: %v", err)
			return
		}

		var result string
		if apiKey != "" {
			// Validate the provided API key
			if err := validations.ValidateString(apiKey); err != nil {
				fmt.Printf("Invalid API key: %v\n", err)
				return
			}
			result = apiKey
			fmt.Print("Using CLI argument for API key\n")
		} else {
			// Fall back to wizard
			fmt.Print("Using interactive prompt for API key\n")
			prompt := promptui.Prompt{
				Label:    "goldapi.io API Key",
				Validate: validations.ValidateString,
				Default:  loadedConfig.RuntimeFlags.GoldAPIKey,
			}

			result, err = prompt.Run()
			if err != nil {
				fmt.Printf("%q failed: %v", "goldapi.io API Key", err)
				return
			}
		}

		loadedConfig.RuntimeFlags.SetGoldAPIKey(result)
	},
}

func init() {
	initCmd.Flags().StringP("api-key", "k", "", "goldapi.io API key to bypass interactive prompt")
	rootCmd.AddCommand(initCmd)
}
