package common

import (
	"fmt"
	"os"
	"strconv"
)

var cfg serviceConfig

type serviceConfig struct {
	name string
	port int
}

type service struct {
	config serviceConfig
}

func (sc *serviceConfig) load() error {
	name := os.Getenv("SERVICE_NAME")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("%s: invalid value for PORT (%d): %w", name, port, err)
	}

	sc.name = name
	sc.port = port

	return nil
}
