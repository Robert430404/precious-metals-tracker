package repositories

import (
	"database/sql"

	"github.com/robert430404/precious-metals-tracker/db"
	"github.com/robert430404/precious-metals-tracker/db/entities"
)

type HoldingRepository struct {
	dbConnection *sql.DB
}

var holdingRepositoryInstance *HoldingRepository = nil

func GetHoldingRepository() *HoldingRepository {
	if holdingRepositoryInstance != nil {
		return holdingRepositoryInstance
	}

	connection, _ := db.GetConnection()

	holdingRepositoryInstance = &HoldingRepository{
		dbConnection: connection,
	}

	return holdingRepositoryInstance
}

func (self *HoldingRepository) GetAllHoldings() []*entities.Holding {
	rows, err := self.dbConnection.Query("select * from holdings")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var holdings []*entities.Holding

	for rows.Next() {
		var alb entities.Holding
		if err := rows.Scan(
			&alb.ID,
			&alb.CreatedAt,
			&alb.UpdatedAt,
			&alb.DeletedAt, 
			&alb.Name, 
			&alb.Source, 
			&alb.PurchaseSpotPrice, 
			&alb.TotalUnits, 
			&alb.UnitWeight,
			&alb.Type, 
		); err != nil {
			break
		}
		holdings = append(holdings, &alb)
	}

	return holdings
}

func (self *HoldingRepository) DeleteHolding(id string) {}

func (self *HoldingRepository) CreateHolding(holding *entities.Holding) {}
