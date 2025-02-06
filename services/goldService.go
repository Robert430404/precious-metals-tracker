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
	var holdings []entities.Holding

	// found := self.connection.Find(&holdings, "type = ?", models.Gold)
	found := self.holdingRepo.GetAllHoldings()
	if found == nil {
		return 0, errors.New("no holdings are present, please add some.")
	}

	totalWeight := self.calculation.CalculateMetalWeight(holdings)

	return totalWeight, nil
}

func (self *GoldService) GetTotalGoldValue() (float64, error) {
	// found := self.connection.Find(&holdings, "type = ?", models.Gold)
	found := self.holdingRepo.GetAllHoldings()
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
