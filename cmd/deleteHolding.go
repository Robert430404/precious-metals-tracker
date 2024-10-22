/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/validations"
	"github.com/spf13/cobra"
)

var deleteHoldingCmd = &cobra.Command{
	Use:   "deleteHolding",
	Short: "Deletes a holding",
	Long:  `Deletes a holding from the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label:    "Holding ID",
			Validate: validations.ValidateTotal,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("%q failed: %v", "Holding ID", err)
			return
		}

		db := db.GetConnection()

		db.Delete(&entities.Holding{}, result)

		fmt.Printf("holding deleted: %v \n", result)
	},
}

func init() {
	rootCmd.AddCommand(deleteHoldingCmd)
}
