package grpc

type ServerConfig struct {
	Port       string `env:"GRPC_PORT,default=50051"`
	Reflection bool   `env:"GRPC_REFLECTION,default=false"`
}
