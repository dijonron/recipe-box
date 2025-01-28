package main

import (
	"log"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
	"github.com/dijonron/recipe-box/server/internal/common"
	"google.golang.org/grpc"
)

func main() {
	service := common.CreateService()

	_, err := service.Serve(func(server *grpc.Server) {
		ingredientpb.RegisterIngredientServiceServer(server, &IngredientServer{})
	})

	if err != nil {
		log.Fatal(err)
	}

}
