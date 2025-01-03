package services

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/robert430404/precious-metals-tracker/renderers"
	"github.com/robert430404/precious-metals-tracker/transformers"
	"github.com/robert430404/precious-metals-tracker/validations"
)

type HoldingService struct {
	LoadedConfig  *config.Config
	TableRenderer renderers.Renderer
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

	hydratedService = &HoldingService{
		LoadedConfig:  config,
		TableRenderer: renderer,
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

	self.TableRenderer.RenderHoldingList(holdings)
}

func (self *HoldingService) GetValue() {
	repository, err := pricing.GetPricingRepository()
	if err != nil {
		fmt.Printf("there was a problem resolving the price repository: %v", err)
		return
	}

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

	var totalWeight float64 = 0
	for _, holding := range holdings {
		unitWeight, err := strconv.ParseFloat(holding.UnitWeight, 16)
		if err != nil {
			continue
		}

		totalUnits, err2 := strconv.ParseFloat(holding.TotalUnits, 16)
		if err2 != nil {
			continue
		}

		totalWeight += unitWeight * totalUnits
	}

	spotPrice := repository.GetSilverSpot()
	totalValue := totalWeight * spotPrice

	self.TableRenderer.RenderValueTable(fmt.Sprintf("$%.2f", totalValue), fmt.Sprintf("$%.2f", spotPrice))
}

func (self *HoldingService) handleAddHoldingFirstRun() error {
	if self.LoadedConfig.RuntimeFlags.AddHoldingRan {
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

	self.LoadedConfig.RuntimeFlags.SetAddHoldingRan(true)

	return nil
}
