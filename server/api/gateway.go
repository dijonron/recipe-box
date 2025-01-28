package main

import (
	"context"
	"log"
	"net/http"
	"os"

	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// service := common.CreateService()

	// // Start the service and register the specific gRPC service
	// _, err := service.Serve(func(server *grpc.Server) {
	// 	recipepb.RegisterRecipeServiceServer(server, &recipe_service.RecipeServer{})
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	log.Println("Starting HTTP server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// gwServer := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// }

	// log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	// log.Fatalln(gwServer.ListenAndServe())
}
