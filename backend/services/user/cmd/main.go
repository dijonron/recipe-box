package user

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dijonron/recipe-box/pkg/database"
	"github.com/dijonron/recipe-box/pkg/grpc"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := GetConfig()
	if cfg.Env == "local" {
		slog.Info("loaded config", "config: ", cfg)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := database.ConnectToDB(ctx, cfg.DatabaseConfig)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	server := buildServer(cfg, db)

	go handleShutdown(ch, ctx, cancel, server)

	// startServer(ctx, httpServer)

	// slog.Info("application terminated")
}

func buildServer(cfg Config, db *sqlx.DB) grpc.Server {

	// userPersistence := up.NewUserPersistence(db)
	// user := user.NewUserService(userPersistence)

	server := grpc.NewServer(cfg.ServerConfig)

	return server
}

func handleShutdown(ch <-chan os.Signal, ctx context.Context, cancel context.CancelFunc, server grpc.Server) {
	<-ch
	server.Stop(ctx)
	cancel()
}

func startServer(ctx context.Context, httpServer grpc.Server) {
	go httpServer.Start(ctx)

	<-ctx.Done()
}