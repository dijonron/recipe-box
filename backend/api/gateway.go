package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
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
	ingreientServiceURL := os.Getenv("INGREDIENT_SERVICE_URL")

	if recipeServiceURL == "" || ingreientServiceURL == "" {
		log.Fatalf("Missing environment variables for service URLs")
	}

	connRecipe, err := grpc.NewClient(recipeServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to recipe service: %v", err)
	}

	connIngredient, err := grpc.NewClient(ingreientServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to ingredient service: %v", err)
	}

	gwmux := runtime.NewServeMux()

	err = recipepb.RegisterRecipeServiceHandler(context.Background(), gwmux, connRecipe)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = ingredientpb.RegisterIngredientServiceHandler(context.Background(), gwmux, connIngredient)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: withCORS(gwmux),
	}

	log.Printf("Serving gRPC-Gateway on port %d ...", port)
	log.Fatalln(gwServer.ListenAndServe())
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
