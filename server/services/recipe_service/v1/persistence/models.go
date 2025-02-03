package recipe_persistence

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Enum representation for Measurement
type Measurement int32

const (
	MeasurementUnspecified Measurement = iota
	MeasurementUnit
	MeasurementTsp
	MeasurementTbsp
	MeasurementCup
	MeasurementMilliliter
	MeasurementLiter
	MeasurementGram
)

type Recipe struct {
	Id          uuid.UUID          `db:"id"`
	Name        string             `db:"name"`
	Chef        string             `db:"chef"`
	Cookbook    string             `db:"Cookbook"`
	Ingredients []RecipeIngredient `db:"ingredients"`
	CreatedAt   time.Time          `db:"created_at"`
	UpdatedAt   time.Time          `db:"updated_at"`
	DeletedAt   sql.NullTime       `db:"deleted_at"`
}

type RecipeIngredient struct {
	IngredientId uuid.UUID    `db:"ingredient_id"`
	RecipeID     uuid.UUID    `db:"recipe_id"`
	Amount       float64      `db:"amount"`
	Measurement  Measurement  `db:"measurement"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}
