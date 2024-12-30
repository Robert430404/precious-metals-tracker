package db

import (
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/robert430404/precious-metals-tracker/config"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB = nil

func GetConnection() (*gorm.DB, error) {
	if dbConnection != nil {
		return dbConnection, nil
	}

	loadedConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	sqlitePath := loadedConfig.SqlitePath

	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, errors.New("could not establish database connection, run init first")
	}

	dbConnection = db

	return dbConnection, nil
}