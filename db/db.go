package db

import (
	"fmt"

	"github.com/robert430404/precious-metals-tracker/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB = nil

func GetConnection() *gorm.DB {
	if dbConnection != nil {
		return dbConnection
	}

	loadedConfig := config.GetConfig()
	sqlitePath := loadedConfig.SqlitePath

	fmt.Print("opening connection to the sqlite database \n")
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		panic("could not establish database connection, run init first")
	}

	dbConnection = db

	return dbConnection
}