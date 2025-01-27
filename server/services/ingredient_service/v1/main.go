package main

import (
	"log"

	"github.com/dijonron/recipe-box/generated/proto/ingredient_service/v1/ingredientpb"
	"github.com/dijonron/recipe-box/internal/common"
	"google.golang.org/grpc"
)

type IngredientServer struct {
	ingredientpb.UnimplementedIngredientServiceServer
}

func main() {
	service := common.CreateService()

	// Start the service and register the specific gRPC service
	_, err := service.Serve(func(server *grpc.Server) {
		// Register the ingredient_service gRPC server implementation
		ingredientpb.RegisterIngredientServiceServer(server, &IngredientServer{})
	})

	if err != nil {
		log.Fatal(err)
	}

}
