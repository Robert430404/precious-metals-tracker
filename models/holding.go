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
	ProductName       string      `json:"product_name"`
	Price             string      `json:"price"`
	Source            string      `json:"source"`
	PurchaseSpotPrice string      `json:"purchase_spot_price"`
	TotalUnits        string      `json:"total_units"`
	UnitWeight        string      `json:"unit_weight"`
	Type              HoldingType `json:"type"`
}

type Prompt struct {
	Label     string
	Validator promptui.ValidateFunc
	Setter    func(value string)
}

func (self *Holding) Hydrate() error {
	fields := [6]Prompt{
		{
			Label:     "Product Name",
			Validator: validations.ValidateString,
			Setter:    func(value string) { self.ProductName = value },
		},
		{
			Label:     "Purchase Price",
			Validator: validations.ValidatePrice,
			Setter:    func(value string) { self.Price = value },
		},
		{
			Label:     "Purchase Source",
			Validator: validations.ValidateString,
			Setter:    func(value string) { self.Source = value },
		},
		{
			Label:     "Spot Price (at time of purchase)",
			Validator: validations.ValidatePrice,
			Setter:    func(value string) { self.PurchaseSpotPrice = value },
		},
		{
			Label:     "How Many Units",
			Validator: validations.ValidateTotal,
			Setter:    func(value string) { self.TotalUnits = value },
		},
		{
			Label:     "Weight Of A Single Unit (in toz)",
			Validator: validations.ValidatePrice,
			Setter:    func(value string) { self.UnitWeight = value },
		},
	}

	for i := 0; i < len(fields); i++ {
		prompt := fields[i]
		value, err := self.PromptForValue(prompt.Label, prompt.Validator)
		if err != nil {
			return err
		}

		prompt.Setter(value)
	}

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
