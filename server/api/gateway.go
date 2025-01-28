package gateway

import (
	"context"
	"log"
	"net/http"
	"os"

	recipepb "github.com/dijonron/recipe-box/generated/recipe_service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGateway() {
	recipeServiceURL := os.Getenv("RECIPE_SERVICE_URL")
	if recipeServiceURL == "" {
		log.Fatalf("Missing environment variables for service URLs")
	}

	connRecipe, err := grpc.NewClient(recipeServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to recipe service: %v", err)
	}
	defer connRecipe.Close()

	mux := runtime.NewServeMux()

	if err := recipepb.RegisterRecipeServiceHandlerServer(context.Background(), mux, recipepb.NewRecipeServiceClient(connRecipe)); err != nil {
		log.Fatalf("could not register recipe service handler: %v", err)
	}

	log.Printf("starting gateway server on port %d ...", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start gateway server on port %d: %v", 8080, err)
	}
}
