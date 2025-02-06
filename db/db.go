package db

import (
	"database/sql"
	"errors"

	"github.com/robert430404/precious-metals-tracker/config"

	_ "github.com/glebarez/sqlite"
)

var dbConnection *sql.DB = nil

func GetConnection() (*sql.DB, error) {
	if dbConnection != nil {
		return dbConnection, nil
	}

	loadedConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	sqlitePath := loadedConfig.SqlitePath

	db, err := sql.Open("sqlite", sqlitePath)
	if err != nil {
		return nil, errors.New("could not establish database connection, run init first")
	}

	dbConnection = db

	return dbConnection, nil
}
