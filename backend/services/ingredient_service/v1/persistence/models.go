package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	Id        uuid.UUID    `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
