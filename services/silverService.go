package services

import (
	"errors"

	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
	"github.com/robert430404/precious-metals-tracker/models"
	"gorm.io/gorm"
)

type SilverService struct {
	connection  *gorm.DB
	repository  *pricing.PricingRepository
	calculation *CalculationService
}

var silverServiceInstance *SilverService = nil

func GetSilverService() (*SilverService, error) {
	if silverServiceInstance != nil {
		return silverServiceInstance, nil
	}

	dbConnection, err := db.GetConnection()
	if err != nil {
		return nil, errors.New("could not get database connection when getting silver service")
	}

	pricingRepo, err := pricing.GetPricingRepository()
	if err != nil {
		return nil, errors.New("could not get pricing repository")
	}

	silverServiceInstance = &SilverService{
		connection:  dbConnection,
		repository:  pricingRepo,
		calculation: GetCalculationService(),
	}

	return silverServiceInstance, nil
}

func (self *SilverService) GetCurrentSilverSpot() float64 {
	return self.repository.GetSilverSpot()
}

func (self *SilverService) GetTotalSilverWeight() (float64, error) {
	var holdings []entities.Holding

	found := self.connection.Find(&holdings, "type = ?", "Silver")
	if found.RowsAffected < 1 {
		return 0, errors.New("no holdings are present, please add some.")
	}

	totalWeight := self.calculation.CalculateMetalWeight(holdings)

	return totalWeight, nil
}

func (self *SilverService) GetTotalSilverValue() (float64, error) {
	var holdings []entities.Holding

	found := self.connection.Find(&holdings, "type = ?", models.Silver)
	if found.RowsAffected < 1 {
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
