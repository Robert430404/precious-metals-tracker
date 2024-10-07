package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/validations"
	"github.com/robert430404/precious-metals-tracker/entities"
	"github.com/spf13/cobra"
)

func PromptForValue(label string, validation promptui.ValidateFunc) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validation,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("%q failed: %v", label, err)
		return result, err
	}

	return result, nil
}

func PromptForType() (string, error) {
	typeSelect := promptui.Select{
		Label: "Holding Type",
		Items: []string{entities.Silver, entities.Gold},
	}

	_, result, err := typeSelect.Run()
	if err != nil {
		fmt.Printf("Invalid type selected %v\n", err)
		return result, err
	}

	return result, nil
}

func HandleAddHolding(cmd *cobra.Command, args []string) {
	purchasePrice, err := PromptForValue("Purchase Price", validations.ValidatePrice)
	if err != nil {
		return
	}

	purchaseSource, err1 := PromptForValue("Purchase Source", validations.ValidateString)
	if err1 != nil {
		return
	}

	spotPrice, err2 := PromptForValue("Spot Price (at time of purchase)", validations.ValidatePrice)
	if err2 != nil {
		return
	}

	totalUnits, err3 := PromptForValue("How Many Units", validations.ValidateTotal)
	if err3 != nil {
		return
	}

	unitWeight, err4 := PromptForValue("Weight Of A Single Unit (in toz)", validations.ValidatePrice)
	if err4 != nil {
		return
	}

	holdingType, err5 := PromptForType()
	if err5 != nil {
		return
	}

	fmt.Printf("addHolding called %q, %q, %q, %q, %q, %q\n", purchasePrice, purchaseSource, spotPrice, totalUnits, unitWeight, holdingType)
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
