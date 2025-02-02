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
	defer connRecipe.Close()

	connIngredient, err := grpc.NewClient(ingreientServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to recipe service: %v", err)
	}
	defer connIngredient.Close()

	mux := runtime.NewServeMux()
	ctx := context.Background()

	if err := recipepb.RegisterRecipeServiceHandlerFromEndpoint(ctx, mux, recipeServiceURL, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	if err := ingredientpb.RegisterIngredientServiceHandlerFromEndpoint(ctx, mux, ingreientServiceURL, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	http.Handle("/", withCORS(mux))
	log.Printf("Starting HTTP server on port %d ...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
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
