package transformers

import (
	"testing"

	"github.com/robert430404/precious-metals-tracker/models"
)

func TestTransformModelToEntity(t *testing.T) {
	transformer := &HoldingTransformer{}

	result := transformer.TransformModelToEntity(&models.Holding{
		ProductName:       "name",
		Price:             "price",
		Source:            "source",
		PurchaseSpotPrice: "spot",
		TotalUnits:        "total",
		UnitWeight:        "weight",
		Type:              "silver",
	})

	if result.Name != "name" {
		t.Fatal("the name did not properly map")
	}

	if result.Source != "source" {
		t.Fatal("the source did not properly map")
	}

	if result.PurchaseSpotPrice != "spot" {
		t.Fatal("the name did not properly map")
	}

	if result.TotalUnits != "total" {
		t.Fatal("the total units did not properly map")
	}

	if result.UnitWeight != "weight" {
		t.Fatal("the unit weight did not properly map")
	}

	if result.Type != "silver" {
		t.Fatal("the type did not properly map")
	}
}
