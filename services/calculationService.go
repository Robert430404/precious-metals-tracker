package services

import (
	"strconv"

	"github.com/robert430404/precious-metals-tracker/db/entities"
)

type CalculationService struct{}

var calculationServiceInstance *CalculationService = nil

func GetCalculationService() *CalculationService {
	if calculationServiceInstance != nil {
		return calculationServiceInstance
	}

	calculationServiceInstance = &CalculationService{}

	return calculationServiceInstance
}

func (self *CalculationService) CalculateMetalWeight(holdings []entities.Holding) float64 {
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

	return totalWeight
}
