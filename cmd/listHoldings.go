package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listHoldingsCmd = &cobra.Command{
	Use:   "listHoldings",
	Short: "Lists all holdings stored in the system",
	Long: `This command lists all preciou metal holdings stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listHoldings called")
	},
}

func init() {
	rootCmd.AddCommand(listHoldingsCmd)
}
