package config

import (
	"log"
	"log/slog"

	"github.com/Netflix/go-env"
	"github.com/dijonron/recipe-box/pkg/grpc"
	"github.com/dijonron/recipe-box/pkg/logger"
)

type AuthConfig struct {
	JWTSecret     string `env:"JWT_SECRET"`
	JWTExpiration int    `env:"JWT_EXPIRATION"`
}

type UserClientConfig struct {
	Host string `env:"USER_SERVICE_HOST,default=localhost"`
	Port string `env:"USER_SERVICE_PORT,default=50051"`
}

type Config struct {
	Env              string `env:"ENV,default=local"`
	ServiceName      string `env:"SERVICE_NAME,default=user"`
	ServerConfig     grpc.ServerConfig
	LoggerConfig     logger.LoggerConfig
	AuthConfig       AuthConfig
	UserClientConfig UserClientConfig
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
		slog.Any("user", c.UserClientConfig),
	)
}
