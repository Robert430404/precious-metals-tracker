package migrations

const SelectMigrationsTableCheckQuery = "select * from migrations"

const SelectMigrationByNameQuery = "select * from migrations where name = ?"

const CreateMigrationsTableQuery = `
create table if not exists migrations (
	id integer primary key autoincrement,
	created_at datetime,
	name text,
	was_completed boolean,
	failure_reason text
)`

const MigrationInsertQuery = `
insert into migrations (
	created_at, 
	name, 
	was_completed, 
	failure_reason
) values (
	?,
	?,
	?,
	?
)`
