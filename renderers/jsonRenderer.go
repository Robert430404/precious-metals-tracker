package renderers

import (
	"encoding/json"
	"fmt"

	"github.com/robert430404/precious-metals-tracker/db/entities"
	"github.com/robert430404/precious-metals-tracker/models"
)

type JsonRenderer struct{}

func (self *JsonRenderer) RenderHoldingList(holdings []entities.Holding) {
	blob, _ := json.Marshal(holdings)

	fmt.Print(string(blob))
}

func (self *JsonRenderer) RenderValueList(data [][]string) {
	jsonString := "["

	for index, value := range data {
		if index > 0 {
			jsonString += ","
		}

		row := "{"

		for index2, value2 := range value {
			key := "type"

			if index2 == 1 {
				key = "currentValue"
			} else if index2 == 2 {
				key = "currentSpotPrice"
			} else if index2 == 3 {
				key = "totalHoldingWeight"
			}

			if index2 > 0 {
				row += ", "
			}

			row += fmt.Sprintf("\"%v\": \"%v\"", key, value2)
		}

		row += "}"

		jsonString += row
	}

	jsonString += "]"
	fmt.Print(jsonString)
}

func (self *JsonRenderer) RenderSpotPricing(silverSpot string, goldSpot string) {
	fmt.Printf(
		"{\"%v\": \"%v\", \"%v\": \"%v\"}",
		models.Silver,
		silverSpot,
		models.Gold,
		goldSpot,
	)
}
