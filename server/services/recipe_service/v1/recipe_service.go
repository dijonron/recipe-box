package main

import (
	"context"
	"fmt"
	"log"
	"os"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	utils "github.com/dijonron/recipe-box/server/internal"
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
	var newIngredients []string                                // ingredients that do not have an id
	var ingredientsToAdd []recipe_persistence.RecipeIngredient // ids of existing ingredients

	for _, ingredient := range ingredients {
		if ingredient.IngredientId == "" {
			newIngredients = append(newIngredients, ingredient.Name)
		}
	}

	for _, ingredient := range ingredients {
		if ingredient.IngredientId != "" {
			id, err := uuid.Parse(ingredient.IngredientId)
			if err != nil {
				log.Fatalf("Failed to parse UUID: %v", err)
			}
			ingredientsToAdd = append(ingredientsToAdd, recipe_persistence.RecipeIngredient{IngredientId: id, Amount: float64(ingredient.Amount), Measurement: recipe_persistence.Measurement(ingredient.Measurement)})
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
			id, err := uuid.Parse(ingredient.Id)
			if err != nil {
				log.Fatalf("Failed to parse UUID: %v", err)
			}

			matchIng := utils.Filter(ingredients, func(s *recipepb.RecipeIngredient) bool { return s.Name == ingredient.Name })
			amount := matchIng[0].Amount
			measurement := matchIng[0].Measurement
			ingredientsToAdd = append(ingredientsToAdd, recipe_persistence.RecipeIngredient{IngredientId: id, Amount: float64(amount), Measurement: recipe_persistence.Measurement(measurement)})
		}
	}

	// save recipe ingredients
	var recipeIngredints []recipe_persistence.RecipeIngredient
	for _, ing := range ingredientsToAdd {

		recId, err := uuid.Parse(saved.Id)
		if err != nil {
			log.Fatalf("Failed to parse UUID: %v", err)
		}
		recipeIngredints = append(recipeIngredints, recipe_persistence.RecipeIngredient{RecipeID: recId, IngredientId: ing.IngredientId, Amount: ing.Amount, Measurement: ing.Measurement})
	}

	err = recipe_persistence.WriteRecipeIngredients(recipeIngredints, s.db)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s RecipeServer) ListRecipes(ctx context.Context, req *recipepb.ListRecipesRequest) (*recipepb.ListRecipesResponse, error) {
	var recipes []recipe_persistence.Recipe

	err := s.db.Select(&recipes, "SELECT id, name, chef FROM rs_recipes WHERE deleted_at IS NULL ORDER BY name ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ingredients: %w", err)
	}

	response := &recipepb.ListRecipesResponse{
		Recipes: make([]*recipepb.Recipe, len(recipes)),
	}

	for i, recipe := range recipes {
		response.Recipes[i] = &recipepb.Recipe{
			Id:   recipe.Id.String(),
			Name: recipe.Name,
			Chef: recipe.Chef,
		}
	}

	return response, nil
}

func (s RecipeServer) DeleteRecipe(ctx context.Context, req *recipepb.DeleteRecipeRequest) (*recipepb.DeleteRecipeResponse, error) {
	recipeId := req.Id

	tx, err := s.db.Beginx()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// delete recipe ingredients
	err = recipe_persistence.DeleteRecipeIngredients(recipeId, tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete recipe ingredients")
	}

	// delete recipe
	err = recipe_persistence.DeleteRecipe(recipeId, tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete recipe")
	}

	return nil, nil
}

func (s RecipeServer) UpdateRecipe(ctx context.Context, req *recipepb.UpdateRecipeRequest) (*recipepb.UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
