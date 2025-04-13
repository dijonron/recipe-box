package grpc

import (
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(hostName, port string) *grpc.ClientConn {
	slog.Info("connecting to grpc server...")
	url := fmt.Sprintf("%s:%s", hostName, port)

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(url, dialOpts...)
	if err != nil {
		slog.Error("failed to connect to grpc server", "error", err)
		return nil
	}

	slog.Info("connected to grpc server", "url", url)
	return conn
}
