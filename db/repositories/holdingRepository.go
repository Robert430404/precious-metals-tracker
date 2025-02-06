package repositories

import (
	"database/sql"
	"fmt"

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
	result, err := self.dbConnection.Query("select * from holdings");
	if err != nil {
		return nil
	}

	fmt.Printf("rows: %v", result)

	return nil
}

func (self *HoldingRepository) DeleteHolding() {}

func (self *HoldingRepository) CreateHolding(holding *entities.Holding) {}