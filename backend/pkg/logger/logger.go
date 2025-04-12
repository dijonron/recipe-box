package logger

import (
	"log/slog"
	"os"
)

func NewLogger(cfg LoggerConfig) *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLevel(cfg.LogLevel),
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

func getLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
