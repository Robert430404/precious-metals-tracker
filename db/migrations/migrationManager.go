package migrations

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/robert430404/precious-metals-tracker/db"
)

type MigrationsManager struct {
	dbConnection *sql.DB
}

type Migration struct {
	ID            uint
	CreatedAt     string
	Name          string
	WasCompleted  bool
	FailureReason string
}

var migrationsManagerInstance *MigrationsManager

func GetMigrationsManager() *MigrationsManager {
	if migrationsManagerInstance != nil {
		return migrationsManagerInstance
	}

	connection, _ := db.GetConnection()

	migrationsManagerInstance = &MigrationsManager{
		dbConnection: connection,
	}

	return migrationsManagerInstance
}

func (self *MigrationsManager) Init() error {
	_, err := self.CreateMigrationsTable()
	if err != nil {
		fmt.Printf("could not create migrations table %v", err)
	}

	self.CreateHoldingTable()

	return err
}

func (self *MigrationsManager) CreateMigrationsTable() (sql.Result, error) {
	_, err := self.dbConnection.Exec(SelectMigrationsTableCheckQuery)
	if err == nil {
		return nil, nil
	}

	return self.dbConnection.Exec(CreateMigrationsTableQuery)
}

func (self *MigrationsManager) CreateHoldingTable() (sql.Result, error) {
	err := self.performMigration(CreateHoldingTableMigrationName, func() error {
		_, err := self.dbConnection.Exec(CreateHoldingTableQuery)

		return err
	})

	return nil, err
}

func (self *MigrationsManager) performMigration(migrationName string, migrationFunc func() error) error {
	rows, err := self.dbConnection.Query(SelectMigrationByNameQuery, migrationName)
	if err != nil {
		return err
	}
	defer rows.Close()

	migrationRows := self.hydrateMigrationRows(rows)

	if len(migrationRows) == 1 && migrationRows[0].WasCompleted {
		return nil
	}

	statement, err := self.dbConnection.Prepare(MigrationInsertQuery)

	if err != nil {
		return err
	}

	err = migrationFunc()
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		time.Now().Format(time.RFC3339),
		migrationName,
		true,
		"was finished",
	)

	return nil
}

func (self *MigrationsManager) hydrateMigrationRows(rows *sql.Rows) []*Migration {
	var migrations []*Migration

	for rows.Next() {
		var alb Migration

		if err := rows.Scan(
			&alb.ID,
			&alb.CreatedAt,
			&alb.Name,
			&alb.WasCompleted,
			&alb.FailureReason,
		); err != nil {
			break
		}

		migrations = append(migrations, &alb)
	}

	return migrations
}
