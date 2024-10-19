package entities

import (
	"github.com/robert430404/precious-metals-tracker/models"
	"gorm.io/gorm"
)

type Holding struct {
	gorm.Model

	Name              string
	Source            string
	PurchaseSpotPrice string
	TotalUnits        string
	UnitWeight        string
	Type              models.HoldingType
}
