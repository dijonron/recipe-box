package main

import (
	"context"
	"log"
	"os"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	recipe_persistence "github.com/dijonron/recipe-box/server/services/recipe_service/v1/persistence"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type RecipeServer struct {
	recipepb.UnimplementedRecipeServiceServer
	db *sqlx.DB
}

func (s RecipeServer) CreateRecipe(ctx context.Context, req *recipepb.CreateRecipeRequest) (*recipepb.CreateRecipeResponse, error) {
	// insert recipe
	recipe := recipe_persistence.Recipe{
		Name:     req.Recipe.GetName(),
		Chef:     req.Recipe.GetChef(),
		Cookbook: req.Recipe.GetCookbook(),
	}

	saved, err := recipe_persistence.WriteRecipe(recipe, s.db)
	if err != nil {
		return nil, err
	}

	ingredients := req.GetRecipe().Ingredients
	var newIngredients []string   // ingredients that do not have an id
	var ingredientsToAdd []string // ids of existing ingredients

	for _, ingredient := range ingredients {
		if ingredient.IngredientId == "" {
			newIngredients = append(newIngredients, ingredient.Name)
		}
	}

	for _, ingredient := range ingredients {
		if ingredient.IngredientId != "" {
			ingredientsToAdd = append(ingredientsToAdd, ingredient.IngredientId)
		}
	}

	// insert any new ingredients
	if len(newIngredients) != 0 {
		connIngredient, err := grpc.NewClient(os.Getenv("INGREDIENT_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to ingredient service: %v", err)
		}
		defer connIngredient.Close()

		ingredientClient := ingredientpb.NewIngredientServiceClient(connIngredient)

		ingReq := &ingredientpb.CreateIngredientsForRecipeRequest{Ingredients: make([]*ingredientpb.Ingredient, len(newIngredients))}
		for i, newIng := range newIngredients {
			ingReq.Ingredients[i] = &ingredientpb.Ingredient{Name: newIng}
		}

		saved, err := ingredientClient.CreateIngredientsForRecipe(context.Background(), ingReq)
		if err != nil {
			log.Fatalf("could not save new ingredients: %v", err)
		}
		for _, ingredient := range saved.Ingredients {
			ingredientsToAdd = append(ingredientsToAdd, ingredient.Id)
		}
	}

	// save recipe ingredients
	var recipeIngredints []recipe_persistence.RecipeIngredient
	for _, ing := range ingredientsToAdd {
		ingId, err := uuid.Parse(ing)
		if err != nil {
			log.Fatalf("Failed to parse UUID: %v", err)
		}
		recId, err := uuid.Parse(saved.Id)
		if err != nil {
			log.Fatalf("Failed to parse UUID: %v", err)
		}
		recipeIngredints = append(recipeIngredints, recipe_persistence.RecipeIngredient{IngredientId: ingId, RecipeID: recId})
	}

	err = recipe_persistence.WriteRecipeIngredients(recipeIngredints, s.db)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s RecipeServer) GetRecipe(ctx context.Context, req *recipepb.GetRecipeRequest) (*recipepb.GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}

func (s RecipeServer) ListRecipes(ctx context.Context, req *recipepb.ListRecipesRequest) (*recipepb.ListRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipes not implemented")
}
func (s RecipeServer) UpdateRecipe(ctx context.Context, req *recipepb.UpdateRecipeRequest) (*recipepb.UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
func (s RecipeServer) DeleteRecipe(ctx context.Context, req *recipepb.DeleteRecipeRequest) (*recipepb.DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
