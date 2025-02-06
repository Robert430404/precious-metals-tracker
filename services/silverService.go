package services

import (
	"errors"

	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/db/repositories"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
)

type SilverService struct {
	holdingRepo *repositories.HoldingRepository
	pricingRepo *pricing.PricingRepository
	calculation *CalculationService
}

var silverServiceInstance *SilverService = nil

func GetSilverService() (*SilverService, error) {
	if silverServiceInstance != nil {
		return silverServiceInstance, nil
	}

	pricingRepo, err := pricing.GetPricingRepository()
	if err != nil {
		return nil, errors.New("could not get pricing repository")
	}

	silverServiceInstance = &SilverService{
		holdingRepo: repositories.GetHoldingRepository(),
		pricingRepo: pricingRepo,
		calculation: GetCalculationService(),
	}

	return silverServiceInstance, nil
}

func (self *SilverService) GetCurrentSilverSpot() float64 {
	return self.pricingRepo.GetSilverSpot()
}

func (self *SilverService) GetTotalSilverWeight() (float64, error) {
	found := self.holdingRepo.GetAllSilverHoldings()
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

func (self *SilverService) GetTotalSilverValue() (float64, error) {
	found := self.holdingRepo.GetAllSilverHoldings()
	if found == nil {
		return 0, errors.New("no holdings are present, please add some.")
	}

	totalWeight, err := self.GetTotalSilverWeight()
	if err != nil {
		return 0, err
	}

	spotPrice := self.GetCurrentSilverSpot()
	totalValue := totalWeight * spotPrice

	return totalValue, nil
}
