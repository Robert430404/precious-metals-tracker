package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/robert430404/precious-metals-tracker/config"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/db/repositories"
	"github.com/robert430404/precious-metals-tracker/models"
	"github.com/robert430404/precious-metals-tracker/renderers"
	"github.com/robert430404/precious-metals-tracker/transformers"
	"github.com/robert430404/precious-metals-tracker/types"
	"github.com/robert430404/precious-metals-tracker/validations"
)

type HoldingService struct {
	loadedConfig   *config.Config
	outputRenderer renderers.Renderer
	silver         SilverService
	gold           GoldService
	platinum       PlatinumService
	holdingRepo    *repositories.HoldingRepository
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

	platinumService, err := GetPlatinumService()
	if err != nil {
		return nil, err
	}

	hydratedService = &HoldingService{
		loadedConfig:   config,
		outputRenderer: renderer,
		silver:         *silverService,
		gold:           *goldService,
		platinum:       *platinumService,
		holdingRepo:    repositories.GetHoldingRepository(),
	}

	return hydratedService, nil
}

func (self *HoldingService) Add(flags *types.HoldingFlags) {
	err := self.handleAddHoldingFirstRun()
	if err != nil {
		fmt.Printf("There was a problem adding your holding: %v", err)
		return
	}

	holding := &models.Holding{}

	// Check if all CLI arguments are provided to bypass wizard
	if flags != nil && flags.HasAllAddFields() {
		// Validate the provided values
		if err := validations.ValidateString(flags.Name); err != nil {
			fmt.Printf("Invalid product name: %v\n", err)
			return
		}
		if err := validations.ValidatePrice(flags.Price); err != nil {
			fmt.Printf("Invalid price: %v\n", err)
			return
		}
		if err := validations.ValidateString(flags.Source); err != nil {
			fmt.Printf("Invalid source: %v\n", err)
			return
		}
		if err := validations.ValidatePrice(flags.SpotPrice); err != nil {
			fmt.Printf("Invalid spot price: %v\n", err)
			return
		}
		if err := validations.ValidateTotal(flags.Units); err != nil {
			fmt.Printf("Invalid units: %v\n", err)
			return
		}
		if err := validations.ValidatePrice(flags.Weight); err != nil {
			fmt.Printf("Invalid weight: %v\n", err)
			return
		}

		// Use CLI arguments to populate holding
		holding.ProductName = flags.Name
		holding.Price = flags.Price
		holding.Source = flags.Source
		holding.PurchaseSpotPrice = flags.SpotPrice
		holding.TotalUnits = flags.Units
		holding.UnitWeight = flags.Weight
		holding.Type = flags.Type

		fmt.Print("Using CLI arguments to create holding\n")
	} else {
		// Fall back to wizard
		fmt.Print("Using interactive wizard to create holding\n")
		err := holding.Hydrate()
		if err != nil {
			fmt.Printf("There was a problem with the wizard: %v", err)
			return
		}
	}

	transformer := transformers.HoldingTransformer{}
	transformed := transformer.TransformModelToEntity(holding)

	self.holdingRepo.CreateHolding(&transformed)
	fmt.Print("stored holding \n")
}

func (self *HoldingService) Delete(holdingID string) {
	// Use CLI argument if provided, otherwise prompt
	var result string
	if holdingID != "" {
		// Validate the provided ID
		if err := validations.ValidateTotal(holdingID); err != nil {
			fmt.Printf("Invalid holding ID: %v\n", err)
			return
		}
		result = holdingID
		fmt.Print("Using CLI argument for holding ID\n")
	} else {
		// Fall back to wizard
		fmt.Print("Using interactive prompt for holding ID\n")
		prompt := promptui.Prompt{
			Label:    "Holding ID",
			Validate: validations.ValidateTotal,
		}

		var err error
		result, err = prompt.Run()
		if err != nil {
			fmt.Printf("%q failed: %v", "Holding ID", err)
			return
		}
	}

	self.holdingRepo.DeleteHolding(result)
	fmt.Printf("holding deleted: %v \n", result)
}

func (self *HoldingService) List() {
	found := self.holdingRepo.GetAllHoldings()
	if found == nil {
		fmt.Print("no holdings are present, please add some. \n")
		return
	}

	transformed := []entities.Holding{}

	for _, holding := range found {
		transformed = append(transformed, *holding)
	}

	self.outputRenderer.RenderHoldingList(transformed)
}

func (self *HoldingService) GetValue() {
	silverSpotPrice := self.silver.GetCurrentSilverSpot()
	silverTotalWeight, err := self.silver.GetTotalSilverWeight()
	if err != nil {
		silverTotalWeight = 0
	}

	silverTotalValue, err := self.silver.GetTotalSilverValue()
	if err != nil {
		silverTotalValue = 0
	}

	goldSpotPrice := self.gold.GetCurrentGoldSpot()
	goldTotalWeight, err := self.gold.GetTotalGoldWeight()
	if err != nil {
		goldTotalWeight = 0
	}

	goldTotalValue, err := self.gold.GetTotalGoldValue()
	if err != nil {
		goldTotalValue = 0
	}

	platinumSpotPrice := self.platinum.GetCurrentPlatinumSpot()
	platinumTotalWeight, err := self.platinum.GetTotalPlatinumWeight()
	if err != nil {
		platinumTotalWeight = 0
	}

	platinumTotalValue, err := self.platinum.GetTotalPlatinumValue()
	if err != nil {
		platinumTotalValue = 0
	}

	self.outputRenderer.RenderValueList([][]string{
		{
			models.Silver,
			fmt.Sprintf("$%.2f", silverTotalValue),
			fmt.Sprintf("$%.2f", silverSpotPrice),
			fmt.Sprintf("%.2foz", silverTotalWeight),
		},
		{
			models.Gold,
			fmt.Sprintf("$%.2f", goldTotalValue),
			fmt.Sprintf("$%.2f", goldSpotPrice),
			fmt.Sprintf("%.2foz", goldTotalWeight),
		},
		{
			models.Platinum,
			fmt.Sprintf("$%.2f", platinumTotalValue),
			fmt.Sprintf("$%.2f", platinumSpotPrice),
			fmt.Sprintf("%.2foz", platinumTotalWeight),
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
