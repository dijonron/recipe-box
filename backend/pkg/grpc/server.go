package grpc

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

type server struct {
	grpcServer *grpc.Server
}

func NewServer(cfg ServerConfig) *server {
	return &server{
		grpcServer: grpc.NewServer(),
	}
}

func (s *server) Start(ctx context.Context) {

}

func (s *server) Stop(ctx context.Context) {
	s.grpcServer.GracefulStop()
	slog.Info("Server gracefully shut down")
}
