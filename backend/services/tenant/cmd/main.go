package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dijonron/recipe-box/pkg/logger"
	"github.com/dijonron/recipe-box/services/tenant/cmd/config"
	tenant "github.com/dijonron/recipe-box/services/tenant/internal"
	"github.com/dijonron/recipe-box/services/tenant/internal/service"
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
	tenant := tenant.NewTenantManager()

	server := service.NewTenantServer(cfg.ServerConfig, tenant)
	return server
}

func handleShutdown(ch <-chan os.Signal, cancel context.CancelFunc, server service.Server) {
	<-ch
	server.Stop()
	cancel()
}
