package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("invalid value for PORT (%d): %v", port, err)
	}

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

	if err := recipepb.RegisterRecipeServiceHandler(context.Background(), mux, connRecipe); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	http.Handle("/", mux)
	log.Printf("Starting HTTP server on port %d ...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
