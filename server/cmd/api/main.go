package main

import (
	"fmt"
	"log"
	"net"

	ingpb "github.com/dijonron/recipe-box/generated/proto/ingredient_service/v1"

	"google.golang.org/grpc"
)

var PORT = "50051"

type IngredientServer struct {
	ingpb.UnimplementedIngredientServiceServer
}


func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))

	if err != nil {
		log.Fatalf("failed to listen on port: %s: %v", PORT, err)
	}

	server := grpc.NewServer()

	ingpb.RegisterIngredientServiceServer(server, &IngredientServer{})

	log.Printf("server running on port %s ...", PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}