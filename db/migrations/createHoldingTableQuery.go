package migrations

const CreateHoldingTableMigrationName = "1.CreateHoldingTable"

const CreateHoldingTableQuery = `
create table if not exists holdings (
	id integer primary key autoincrement,
	created_at datetime,
	updated_at datetime,
	deleted_at datetime,
	name text,
	source text,
	purchase_spot_price text,
	total_units text,
	unit_weight text,
	type text
)`
