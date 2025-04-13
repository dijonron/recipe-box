package config

import (
	"log"
	"log/slog"

	"github.com/Netflix/go-env"
	"github.com/dijonron/recipe-box/pkg/grpc"
	"github.com/dijonron/recipe-box/pkg/logger"
)

type Config struct {
	Env          string `env:"ENV,default=local"`
	ServiceName  string `env:"SERVICE_NAME,default=tenant"`
	ServerConfig grpc.ServerConfig
	LoggerConfig logger.LoggerConfig
}

func GetConfig() Config {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
		return Config{}
	}

	return cfg
}

func (c Config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("env", c.Env),
		slog.String("service", c.ServiceName),
		slog.Any("server", c.ServerConfig),
		slog.Any("logger", c.LoggerConfig),
	)
}
