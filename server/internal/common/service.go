package common

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func CreateService() *service {
	err := cfg.load()
	if err != nil {
		log.Fatal(err)
	}
	service := &service{
		config: cfg,
	}

	return service
}

func (s *service) Serve(registerService func(server *grpc.Server)) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.port))

	if err != nil {
		log.Fatalf("%s: failed to listen on port: %d: %v", s.config.name, s.config.port, err)
	}

	server := grpc.NewServer()

	registerService(server)

	reflection.Register(server)
	log.Println("gRPC reflection enabled")

	log.Printf("%s running on port %d ...", s.config.name, s.config.port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return server, nil
}
