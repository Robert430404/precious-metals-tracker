package renderers

import "github.com/robert430404/precious-metals-tracker/db/entities"

type Renderer interface {
	RenderHoldingList(holdings []entities.Holding)
	RenderValueTable(data [][]string)
}
