package user

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/dijonron/recipe-box/pkg/database"
	"github.com/dijonron/recipe-box/pkg/grpc"
)


type Config struct {
	Env            string `env:"ENV,default=local"`
	ServerConfig   grpc.ServerConfig
	DatabaseConfig database.DatabaseConfig
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