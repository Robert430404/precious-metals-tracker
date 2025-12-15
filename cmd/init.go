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
	Use:   "init",
	Short: "Initializes the application in your environment",
	Long:  `This command sets up and migrates the SQLite database for storing your holding information.`,
	Run: func(cmd *cobra.Command, args []string) {
		loadedConfig, err := config.GetConfig()
		if err != nil {
			fmt.Printf("there was a problem loading your configuration: %v", err)
		}

		err2 := os.MkdirAll(loadedConfig.ConfigPath, 0755)
		if err2 != nil {
			fmt.Print("there was a problem creating the configuration directory \n")
			fmt.Printf("ensure that %q exists", loadedConfig.ConfigPath)

			return
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
			return
		}

		prompt := promptui.Prompt{
			Label:    "goldapi.io API Key",
			Validate: validations.ValidateString,
			Default:  loadedConfig.RuntimeFlags.GoldAPIKey,
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
