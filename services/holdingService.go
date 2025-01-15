package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/robert430404/precious-metals-tracker/renderers"
	"github.com/robert430404/precious-metals-tracker/transformers"
	"github.com/robert430404/precious-metals-tracker/validations"
)

type HoldingService struct {
	loadedConfig   *config.Config
	outputRenderer renderers.Renderer
	silver         SilverService
	gold           GoldService
}

var hydratedService *HoldingService = nil

func GetHoldingService(outputType string) (*HoldingService, error) {
	if hydratedService != nil {
		return hydratedService, nil
	}

	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	var renderer renderers.Renderer = nil
	if outputType == "json" {
		renderer = &renderers.JsonRenderer{}
	} else {
		renderer = &renderers.TableRenderer{}
	}

	silverService, err := GetSilverService()
	if err != nil {
		return nil, err
	}

	goldService, err := GetGoldService()
	if err != nil {
		return nil, err
	}

	hydratedService = &HoldingService{
		loadedConfig:   config,
		outputRenderer: renderer,
		silver:         *silverService,
		gold:           *goldService,
	}

	return hydratedService, nil
}

func (self *HoldingService) Add() {
	err := self.handleAddHoldingFirstRun()
	if err != nil {
		fmt.Printf("There was a problem adding your holding: %v", err)
		return
	}

	holding := &models.Holding{}
	holding.Hydrate()

	db, err := db.GetConnection()
	if err != nil {
		fmt.Printf("there was a problem resolving the db connection: %v", err)
		return
	}

	transformer := transformers.HoldingTransformer{}
	transformed := transformer.TransformModelToEntity(holding)

	db.Create(&transformed)
	fmt.Print("stored holding \n")
}

func (self *HoldingService) Delete() {
	prompt := promptui.Prompt{
		Label:    "Holding ID",
		Validate: validations.ValidateTotal,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("%q failed: %v", "Holding ID", err)
		return
	}

	db, err := db.GetConnection()
	if err != nil {
		fmt.Printf("there was a problem resolving the db connection: %v", err)
		return
	}

	db.Delete(&entities.Holding{}, result)

	fmt.Printf("holding deleted: %v \n", result)
}

func (self *HoldingService) List() {
	db, err := db.GetConnection()
	if err != nil {
		fmt.Printf("there was a problem resolving the db connection: %v", err)
		return
	}

	var holdings []entities.Holding

	found := db.Find(&holdings)
	if found.RowsAffected < 1 {
		fmt.Print("no holdings are present, please add some. \n")
		return
	}

	self.outputRenderer.RenderHoldingList(holdings)
}

func (self *HoldingService) GetValue() {
	spotPrice := self.silver.GetCurrentSilverSpot()

	totalWeight, err := self.silver.GetTotalSilverWeight()
	if err != nil {
		fmt.Printf("there was a problem getting total silver weight: %v \n", err)
		return
	}

	totalValue, err := self.silver.GetTotalSilverValue()
	if err != nil {
		fmt.Printf("there was a problem getting total silver value: %v \n", err)
		return
	}
	goldspotPrice := self.gold.GetCurrentGoldSpot()

	goldtotalWeight, err := self.gold.GetTotalGoldWeight()
	if err != nil {
		fmt.Printf("there was a problem getting total gold weight: %v \n", err)
		return
	}

	goldtotalValue, err := self.gold.GetTotalGoldValue()
	if err != nil {
		fmt.Printf("there was a problem getting total gold value: %v \n", err)
		return
	}

	self.outputRenderer.RenderValueTable([][]string{
		{
			models.Silver,
			fmt.Sprintf("$%.2f", totalValue),
			fmt.Sprintf("$%.2f", spotPrice),
			fmt.Sprintf("%.2foz", totalWeight),
		},
		{
			models.Gold,
			fmt.Sprintf("$%.2f", goldtotalValue),
			fmt.Sprintf("$%.2f", goldspotPrice),
			fmt.Sprintf("%.2foz", goldtotalWeight),
		},
	})
}

func (self *HoldingService) handleAddHoldingFirstRun() error {
	if self.loadedConfig.RuntimeFlags.AddHoldingRan {
		return nil
	}

	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	configPath := config.ConfigPath
	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return errors.New(fmt.Sprintf("there was a problem ensuring the config path: %v\n", err))
	}

	self.loadedConfig.RuntimeFlags.SetAddHoldingRan(true)

	return nil
}
