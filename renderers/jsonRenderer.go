package renderers

import (
	"encoding/json"
	"fmt"

	"github.com/robert430404/precious-metals-tracker/db/entities"
)

type JsonRenderer struct{}

func (self *JsonRenderer) RenderHoldingList(holdings []entities.Holding) {
	blob, _ := json.Marshal(holdings)

	fmt.Print(string(blob))
}

func (self *JsonRenderer) RenderValueTable(value string, spotPrice string) {
	fmt.Printf("{\"value\": \"%v\", \"spotPrice\": \"%v\"}", value, spotPrice)
}
