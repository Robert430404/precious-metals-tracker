package entities

import (
	"github.com/robert430404/precious-metals-tracker/models"
)

type Holding struct {
	AttributesEntity

	Name              string
	Source            string
	PurchaseSpotPrice string
	TotalUnits        string
	UnitWeight        string
	Type              models.HoldingType
}
