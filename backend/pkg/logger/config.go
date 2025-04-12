package logger

type LoggerConfig struct {
	LogLevel string `env:"LOG_LEVEL,default=info"`
}
