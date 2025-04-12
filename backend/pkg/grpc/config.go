package grpc

type ServerConfig struct {
	Port string `env:"GRPC_PORT,default=50051"`
	ServiceName string `env:"SERVICE_NAME"`
}
