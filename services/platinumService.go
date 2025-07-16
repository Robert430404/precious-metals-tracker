package services

import (
	"errors"

	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/db/repositories"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
)

type PlatinumService struct {
	holdingRepo *repositories.HoldingRepository
	pricingRepo *pricing.PricingRepository
	calculation *CalculationService
}

var platinumServiceInstance *PlatinumService = nil

func GetPlatinumService() (*PlatinumService, error) {
	if platinumServiceInstance != nil {
		return platinumServiceInstance, nil
	}

	pricingRepo, err := pricing.GetPricingRepository()
	if err != nil {
		return nil, errors.New("could not get pricing repository")
	}

	platinumServiceInstance = &PlatinumService{
		holdingRepo: repositories.GetHoldingRepository(),
		pricingRepo: pricingRepo,
		calculation: GetCalculationService(),
	}

	return platinumServiceInstance, nil
}

func (self *PlatinumService) GetCurrentPlatinumSpot() float64 {
	return self.pricingRepo.GetPlatinumSpot()
}

func (self *PlatinumService) GetTotalPlatinumWeight() (float64, error) {
	found := self.holdingRepo.GetAllPlatinumHoldings()
	if found == nil {
		return 0, errors.New("no holdings are present, please add some.")
	}

	transformed := []entities.Holding{}
	for _, holding := range found {
		transformed = append(transformed, *holding)
	}

	totalWeight := self.calculation.CalculateMetalWeight(transformed)

	return totalWeight, nil
}

func (self *PlatinumService) GetTotalPlatinumValue() (float64, error) {
	found := self.holdingRepo.GetAllPlatinumHoldings()
	if found == nil {
		return 0, errors.New("no holdings are present, please add some.")
	}

	totalWeight, err := self.GetTotalPlatinumWeight()
	if err != nil {
		return 0, err
	}

	spotPrice := self.GetCurrentPlatinumSpot()
	totalValue := totalWeight * spotPrice

	return totalValue, nil
}
