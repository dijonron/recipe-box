package service

import (
	"context"
	"log/slog"

	"github.com/dijonron/recipe-box/pkg/grpc"
)

type Server interface {
	Start(ctx context.Context)
	Stop()
}

type server struct {
	server grpc.GrpcServer
}

var _ Server = (*server)(nil)

func (s server) Start(ctx context.Context) {
	slog.Info("starting server...")
	go func() {
		if err := s.server.Serve(); err != nil {
			slog.Error("failed to start server", "error: ", err)
		}
		slog.Info("server started")
	}()

	<-ctx.Done()
}

func (s server) Stop() {
	s.server.Stop()
}
