package services

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/robert430404/precious-metals-tracker/transformers"
	"github.com/robert430404/precious-metals-tracker/validations"
	"github.com/rodaine/table"
)

type HoldingService struct {
	LoadedConfig *config.Config
}

var hydratedService *HoldingService = nil

func GetHoldingService() *HoldingService {
	if hydratedService != nil {
		return hydratedService
	}

	hydratedService = &HoldingService{
		LoadedConfig: config.GetConfig(),
	}

	return hydratedService
}

func (self *HoldingService) Add() {
	self.handleAddHoldingFirstRun()

	holding := &models.Holding{}
	holding.Hydrate()

	db := db.GetConnection()
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

	db := db.GetConnection()

	db.Delete(&entities.Holding{}, result)

	fmt.Printf("holding deleted: %v \n", result)
}

func (self *HoldingService) List() {
		db := db.GetConnection()

	var holdings []entities.Holding

	found := db.Find(&holdings)
	if found.RowsAffected < 1 {
		fmt.Print("no holdings are present, please add some. \n")
		return
	}

	self.renderHoldingList(holdings)
}

func (self *HoldingService) renderHoldingList(holdings []entities.Holding) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	table := table.New(
		"ID",
		"Name",
		"Purchase Spot Price",
		"Total Units",
		"Unit Weight",
		"Type",
	)

	table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, holding := range holdings {
		table.AddRow(
			holding.ID,
			holding.Name,
			holding.PurchaseSpotPrice,
			holding.TotalUnits,
			holding.UnitWeight,
			holding.Type,
		)
	}

	table.Print()
}

func (self *HoldingService) handleAddHoldingFirstRun() {
	if self.LoadedConfig.RuntimeFlags.AddHoldingRan {
		return
	}

	configPath := config.GetConfig().ConfigPath

	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("there was a problem ensuring the config path: %v\n", err))
	}

	self.LoadedConfig.RuntimeFlags.SetAddHoldingRan(true)
}
