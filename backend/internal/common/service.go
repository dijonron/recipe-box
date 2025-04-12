package common

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
		log.Fatalf("%s: failed to serve: %v", s.config.name, err)
	}

	return server, nil
}

func (s *service) ConnectToDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s password=%s port=%d host=database sslmode=disable", s.config.db_user, s.config.db_name, s.config.db_password, s.config.db_port))
	log.Printf("%s connected to db ...", s.config.name)
	if err != nil {
		log.Fatalf("%s: failed to connect to db: %v", s.config.name, err)
	}

	return db
}
