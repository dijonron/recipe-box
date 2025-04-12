package main

import (
	"log"

	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	"github.com/dijonron/recipe-box/server/internal/common"
	"google.golang.org/grpc"
)

func main() {
	service := common.CreateService()
	db := service.ConnectToDB()

	_, err := service.Serve(func(server *grpc.Server) {
		recipepb.RegisterRecipeServiceServer(server, &RecipeServer{db: db})
	})

	if err != nil {
		log.Fatal(err)
	}

}
