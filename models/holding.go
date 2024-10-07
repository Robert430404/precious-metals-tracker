package models

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/validations"
)

type HoldingType = string

const (
	Silver HoldingType = "Silver"
	Gold   HoldingType = "Gold"
)

type Holding struct {
	Price        string
	Source       string
	PurchaseSpot string
	Volume       string
	Weight       string
	Type         HoldingType
}

func (self *Holding) Hydrate() error {
	purchasePrice, err := self.PromptForValue("Purchase Price", validations.ValidatePrice)
	if err != nil {
		return err
	}

	self.Price = purchasePrice

	purchaseSource, err1 := self.PromptForValue("Purchase Source", validations.ValidateString)
	if err1 != nil {
		return err1
	}

	self.Source = purchaseSource

	spotPrice, err2 := self.PromptForValue("Spot Price (at time of purchase)", validations.ValidatePrice)
	if err2 != nil {
		return err2
	}

	self.PurchaseSpot = spotPrice

	totalUnits, err3 := self.PromptForValue("How Many Units", validations.ValidateTotal)
	if err3 != nil {
		return err3
	}

	self.Volume = totalUnits

	unitWeight, err4 := self.PromptForValue("Weight Of A Single Unit (in toz)", validations.ValidatePrice)
	if err4 != nil {
		return err4
	}

	self.Weight = unitWeight

	holdingType, err5 := self.PromptForType()
	if err5 != nil {
		return err5
	}

	self.Type = holdingType

	return nil
}

func (*Holding) PromptForValue(label string, validation promptui.ValidateFunc) (string, error) {
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

func (*Holding) PromptForType() (string, error) {
	typeSelect := promptui.Select{
		Label: "Holding Type",
		Items: []string{Silver, Gold},
	}

	_, result, err := typeSelect.Run()
	if err != nil {
		fmt.Printf("Invalid type selected %v\n", err)
		return result, err
	}

	return result, nil
}
