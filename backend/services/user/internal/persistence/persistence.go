package persistence

import (
	"context"

	"github.com/dijonron/recipe-box/services/user/internal"
	"github.com/jmoiron/sqlx"
)

type persistence struct {
	*sqlx.DB
}

var _ user.Persistence = (*persistence)(nil)

func NewPersistence(db *sqlx.DB) persistence {
	return persistence{
		DB: db,
	}
}

func (p persistence) SaveUser(ctx context.Context) error {
	// Implementation for saving user to the database
	return nil
}

func (p persistence) GetUserByID(ctx context.Context, userID string) (user.User, error) {
	// Implementation for saving user to the database
	return user.User{}, nil
}

func (p persistence) UpdateUser(ctx context.Context, userID string) error {
	// Implementation for saving user to the database
	return nil
}

func (p persistence) DeleteUser(ctx context.Context, userID string) error {
	// Implementation for saving user to the database
	return nil
}
