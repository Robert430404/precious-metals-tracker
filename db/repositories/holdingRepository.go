package repositories

import (
	"database/sql"
	"time"

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

func (self *HoldingRepository) DeleteHolding(id string) error {
	_, err := self.dbConnection.Exec("delete from holdings where id = ?", id)

	return err
}

func (self *HoldingRepository) CreateHolding(holding *entities.Holding) error {
	statement, err := self.dbConnection.Prepare(`insert into holdings (
		created_at, 
		updated_at, 
		name, 
		source, 
		purchase_spot_price, 
		total_units, 
		unit_weight, 
		type
	) values (
		?, 
		?, 
		?, 
		?, 
		?, 
		?, 
		?, 
		?
	)`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(
		time.Now().Format(time.RFC3339), 
		time.Now().Format(time.RFC3339), 
		holding.Name, 
		holding.Source, 
		holding.PurchaseSpotPrice, 
		holding.TotalUnits, 
		holding.UnitWeight, 
		holding.Type, 
	)

	return err
}
