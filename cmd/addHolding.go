package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addHoldingCmd = &cobra.Command{
	Use:   "addHolding",
	Short: "Adds a precious metals holding.",
	Long:	 `This command walks you through adding a precious metals holding.

It requests the following information:

	- Purchase Price
	- Purchase Source
	- Spot Price at time of purchase
	- Weight of holding
	- Type of holding

This then stores it for use inside of the tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addHolding called")
	},
}

func init() {
	rootCmd.AddCommand(addHoldingCmd)
}
