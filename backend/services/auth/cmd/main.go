package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dijonron/recipe-box/pkg/logger"
	"github.com/dijonron/recipe-box/services/auth/cmd/config"
	auth "github.com/dijonron/recipe-box/services/auth/internal"
	"github.com/dijonron/recipe-box/services/auth/internal/service"
	userclient "github.com/dijonron/recipe-box/services/auth/internal/userclient"
)

func main() {
	cfg := config.GetConfig()
	logger := logger.NewLogger(cfg.LoggerConfig)
	slog.Info(fmt.Sprintf("loaded config for %s service", cfg.ServiceName), "config", cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := buildServer(ctx, cfg)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go handleShutdown(ch, cancel, server)

	server.Start(ctx)

	logger.Info("application terminated")
}

func buildServer(ctx context.Context, cfg config.Config) service.Server {
	userClient := userclient.NewUserClient(cfg.UserClientConfig)
	auth := auth.NewAuthManager(cfg.AuthConfig, userClient)

	server := service.NewAuthServer(cfg.ServerConfig, auth)
	return server
}

func handleShutdown(ch <-chan os.Signal, cancel context.CancelFunc, server service.Server) {
	<-ch
	server.Stop()
	cancel()
}
