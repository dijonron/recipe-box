package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	FAILED_TO_CONNECT = "failed to connect to db"
)

func (cfg DatabaseConfig) GetDBParams() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s port=%s host=%s %s", cfg.User, cfg.Name, cfg.Password, cfg.Port, cfg.Host, cfg.Options)
}

func ConnectToDB(ctx context.Context, cfg DatabaseConfig) (*sqlx.DB, error) {
	params := cfg.GetDBParams()

	db, err := sqlx.ConnectContext(ctx, cfg.Engine, params)
	if err != nil {
		slog.Error(FAILED_TO_CONNECT, "error", err)
		return nil, err
	}
	slog.Info("connected to db")

	return db, nil
}
