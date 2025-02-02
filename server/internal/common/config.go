package common

import (
	"fmt"
	"os"
	"strconv"
)

var cfg serviceConfig

type serviceConfig struct {
	name        string
	port        int
	db_user     string
	db_password string
	db_name     string
	db_port     int
}

type service struct {
	config serviceConfig
}

func (sc *serviceConfig) load() error {
	name := os.Getenv("SERVICE_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return fmt.Errorf("%s: invalid value for DB_PORT (%d): %w", name, db_port, err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("%s: invalid value for PORT (%d): %w", name, port, err)
	}

	sc.name = name
	sc.port = port
	sc.db_user = db_user
	sc.db_password = db_password
	sc.db_name = db_name
	sc.db_port = db_port

	return nil
}
