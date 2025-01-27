package main

import (
	"log"

	"github.com/dijonron/recipe-box/generated/proto/recipe_service/v1/recipepb"
	"github.com/dijonron/recipe-box/internal/common"
	"google.golang.org/grpc"
)

type RecipeServer struct {
	recipepb.UnimplementedRecipeServiceServer
}

func main() {
	service := common.CreateService()

	// Start the service and register the specific gRPC service
	_, err := service.Serve(func(server *grpc.Server) {
		// Register the recipe_service gRPC server implementation
		recipepb.RegisterRecipeServiceServer(server, &RecipeServer{})
	})

	if err != nil {
		log.Fatal(err)
	}

}
