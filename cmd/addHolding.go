package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/validations"
)

func HandleAddHolding(cmd *cobra.Command, args []string) {
	purchasePricePrompt := promptui.Prompt{
		Label:    "Purchase Price",
		Validate: validations.ValidatePrice,
	}

	purchasePrice, err := purchasePricePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	purchaseSourcePrompt := promptui.Prompt{
		Label:    "Purchase Source",
		Validate: validations.ValidateString,
	}

	purchaseSource, err1 := purchaseSourcePrompt.Run()
	if err1 != nil {
		fmt.Printf("Purchase Source Failed %v\n", err1)
		return
	}

	fmt.Printf("addHolding called %q, %q\n", purchasePrice, purchaseSource)
}

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
	Run: HandleAddHolding,
}

func init() {
	rootCmd.AddCommand(addHoldingCmd)
}
