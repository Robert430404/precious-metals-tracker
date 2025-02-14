package entities

import (
	"database/sql"
	"time"
)

type AttributesEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
