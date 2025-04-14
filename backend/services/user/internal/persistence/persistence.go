package persistence

import (
	"context"
	"log/slog"

	user "github.com/dijonron/recipe-box/services/user/internal"
	"github.com/jmoiron/sqlx"
)

const (
	FAILED_TO_GET  string = "failed to fetch user"
	FAILED_TO_SAVE string = "failed to save user"
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

func (p persistence) SaveUser(ctx context.Context, name, email, password string) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`
	_, err := p.ExecContext(ctx, query, name, email, password)
	if err != nil {
		slog.Error(FAILED_TO_SAVE, "err", err)
		return err
	}

	return nil
}

func (p persistence) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	var u User
	query := `SELECT name, email, password_hash, tenant_id, role FROM users WHERE email = $1`
	err := p.GetContext(ctx, &u, query, email)
	if err != nil {
		slog.Error(FAILED_TO_GET, "err", err)
		return user.User{}, err
	}

	return toUser(u), nil
}

func (p persistence) UpdateUserLogin(ctx context.Context, email string) error {
	query := `UPDATE users SET last_login = NOW() WHERE email = $1`
	_, err := p.ExecContext(ctx, query, email)
	if err != nil {
		slog.Error(FAILED_TO_SAVE, "err", err)
	}

	return nil
}

func (p persistence) DeleteUser(ctx context.Context, userID string) error {
	// Implementation for saving user to the database
	return nil
}
