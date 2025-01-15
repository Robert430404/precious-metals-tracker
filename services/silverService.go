package services

import (
	"errors"
	"strconv"

	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/http/pricing"
	"github.com/robert430404/precious-metals-tracker/models"
	"gorm.io/gorm"
)

type SilverService struct {
	connection *gorm.DB
	repository *pricing.PricingRepository
}

var silverServiceInstance *SilverService = nil

func GetSilverService() (*SilverService, error) {
	if silverServiceInstance != nil {
		return silverServiceInstance, nil
	}

	con, err := db.GetConnection()
	if err != nil {
		return nil, errors.New("could got get database connection when getting silver service")
	}

	silverServiceInstance = &SilverService{
		connection: con,
	}

	return silverServiceInstance, nil
}

func (self *SilverService) GetTotalSilverWeight() (float64, error) {
	var holdings []entities.Holding

	found := self.connection.Find(&holdings, "type = ?", "Silver")
	if found.RowsAffected < 1 {
		return 0, errors.New("no holdings are present, please add some.")
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

	return totalWeight, nil
}

func (self *SilverService) GetCurrentSilverSpot() float64 {
	return self.repository.GetSilverSpot()
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
