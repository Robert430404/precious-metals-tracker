package transformers

import (
	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/models"
)

type HoldingTransformer struct{}

func (*HoldingTransformer) TransformModelToEntity(model *models.Holding) entities.Holding {
	return entities.Holding{
		Name:              model.ProductName,
		Source:            model.Source,
		PurchaseSpotPrice: model.PurchaseSpotPrice,
		TotalUnits:        model.TotalUnits,
		UnitWeight:        model.UnitWeight,
		Type:              model.Type,
	}
}
