package database

import (
	"log/slog"
)

type DatabaseConfig struct {
	Engine   string `env:"DB_ENGINE,default=postgres"`
	Host     string `env:"DB_HOST,default=localhost"`
	Name     string `env:"DB_NAME,default=tr_db"`
	Port     string `env:"DB_PORT,default=5432"`
	User     string `env:"DB_USER,default=root"`
	Password string `env:"DB_PASSWORD,default=pwd"`
	Options  string `env:"DB_OPTIONS"`
}

func (c DatabaseConfig) LogValue() slog.Value {
	maskedUser := maskString(c.User)
	maskedPassword := maskString(c.Password)
	return slog.GroupValue(
		slog.String("user", maskedUser),
		slog.String("password", maskedPassword),
	)
}

func maskString(s string) string {
	return "********"
}
