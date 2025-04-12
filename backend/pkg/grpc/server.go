package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer interface {
	Serve() error
	Stop()
}

type Register interface {
	Register(server *grpc.Server)
}

type grpcServer struct {
	server *grpc.Server
	port   string
}

var _ GrpcServer = (*grpcServer)(nil)

func NewServer(cfg ServerConfig, registers []Register) GrpcServer {
	gs := &grpcServer{
		server: grpc.NewServer(),
		port:   cfg.Port,
	}

	for _, r := range registers {
		r.Register(gs.server)
	}

	if cfg.Reflection {
		reflection.Register(gs.server)
	}

	return gs
}

func (g *grpcServer) Serve() error {
	tcp, err := net.Listen("tcp", fmt.Sprintf(":%s", g.port))
	if err != nil {
		slog.Error("failed to listen", "error: ", err)
		return err
	}

	return g.server.Serve(tcp)
}

func (g *grpcServer) Stop() {
	g.server.GracefulStop()
	slog.Info("Server gracefully shut down")
}
