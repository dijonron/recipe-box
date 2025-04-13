package service

import (
	"github.com/dijonron/recipe-box/pkg/grpc"
)

type Config struct {
	GrpcConfig grpc.ServerConfig
}
