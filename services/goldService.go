package services

import (
	"errors"

	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/db/repositories"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
)

type GoldService struct {
	holdingRepo *repositories.HoldingRepository
	pricingRepo *pricing.PricingRepository
	calculation *CalculationService
}

var goldServiceInstance *GoldService = nil

func GetGoldService() (*GoldService, error) {
	if goldServiceInstance != nil {
		return goldServiceInstance, nil
	}

	pricingRepo, err := pricing.GetPricingRepository()
	if err != nil {
		return nil, errors.New("could not get pricing repository")
	}

	goldServiceInstance = &GoldService{
		holdingRepo: repositories.GetHoldingRepository(),
		pricingRepo: pricingRepo,
		calculation: GetCalculationService(),
	}

	return goldServiceInstance, nil
}

func (self *GoldService) GetCurrentGoldSpot() float64 {
	return self.pricingRepo.GetGoldSpot()
}

func (self *GoldService) GetTotalGoldWeight() (float64, error) {
	found := self.holdingRepo.GetAllGoldHoldings()
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

func (self *GoldService) GetTotalGoldValue() (float64, error) {
	found := self.holdingRepo.GetAllGoldHoldings()
	if found == nil {
		return 0, errors.New("no holdings are present, please add some.")
	}

	totalWeight, err := self.GetTotalGoldWeight()
	if err != nil {
		return 0, err
	}

	spotPrice := self.GetCurrentGoldSpot()
	totalValue := totalWeight * spotPrice

	return totalValue, nil
}
