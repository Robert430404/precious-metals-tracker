package cmd

import (
	"fmt"
	"os"

	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the application in your environment",
	Long:  `This command sets up and migrates the SQLite database for storing the holdings`,
	Run: func(cmd *cobra.Command, args []string) {
		loadedConfig := config.GetConfig()
		sqlitePath := loadedConfig.SqlitePath

		_, err := os.Stat(sqlitePath)
		if err != nil {
			fmt.Print("creating sqlite file because it does not exist \n")
			os.Create(sqlitePath)
		}

		fmt.Print("opening connection to the sqlite database \n")
		db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		fmt.Print("migrating the database \n")
		db.AutoMigrate(&entities.Holding{})
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
