package main

import (
	"context"
	"fmt"
	"log"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
	models "github.com/dijonron/recipe-box/server/services/ingredient_service/v1/persistence"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IngredientServer struct {
	ingredientpb.UnimplementedIngredientServiceServer
	db *sqlx.DB
}

func (s IngredientServer) CreateIngredientsForRecipe(ctx context.Context, req *ingredientpb.CreateIngredientsForRecipeRequest) (*ingredientpb.CreateIngredientsForRecipeResponse, error) {
	fmt.Println("Hello!")
	// validate request
	if len(req.GetIngredients()) == 0 {
		return nil, fmt.Errorf("no ingredients provided")
	}

	// map request for sql statement
	ingredientNames := make([]string, 0, len(req.GetIngredients()))
	for _, ingredientReq := range req.GetIngredients() {
		ingredientNames = append(ingredientNames, ingredientReq.GetName())
	}

	// begin and handle txn
	tx, err := s.db.Beginx()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()

		} else {
			err = tx.Commit()
		}
	}()

	// bulk insert ingredient names
	var insertedIngredients []ingredientpb.Ingredient

	query := `
		INSERT INTO is_ingredients (name)
		SELECT * FROM unnest($1::text[]) AS t(name)
		RETURNING id, name
	`

	err = tx.Select(&insertedIngredients, query, pq.Array(ingredientNames))
	if err != nil {
		log.Println("Error inserting ingredients:", err)
		return nil, fmt.Errorf("failed to insert ingredients: %w", err)
	}

	response := &ingredientpb.CreateIngredientsForRecipeResponse{
		Ingredients: make([]*ingredientpb.Ingredient, len(insertedIngredients)),
	}
	for i, ing := range insertedIngredients {
		response.Ingredients[i] = &ingredientpb.Ingredient{
			Id:   ing.Id,
			Name: ing.Name,
		}
	}

	return response, nil
}

func (s IngredientServer) ListIngredients(ctx context.Context, req *ingredientpb.ListIngredientsRequest) (*ingredientpb.ListIngredientsResponse, error) {
	var ingredients []models.Ingredient

	err := s.db.Select(&ingredients, "SELECT id, name FROM is_ingredients WHERE deleted_at IS NULL ORDER BY name ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ingredients: %w", err)
	}

	response := &ingredientpb.ListIngredientsResponse{
		Ingredients: make([]*ingredientpb.Ingredient, len(ingredients)),
	}

	for i, ingredient := range ingredients {
		response.Ingredients[i] = &ingredientpb.Ingredient{
			Id:   ingredient.Id.String(),
			Name: ingredient.Name,
		}
	}

	return response, nil
}

func (s IngredientServer) GetIngredient(ctx context.Context, req *ingredientpb.GetIngredientRequest) (*ingredientpb.GetIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}
func (s IngredientServer) UpdateIngredient(ctx context.Context, req *ingredientpb.UpdateIngredientRequest) (*ingredientpb.UpdateIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}
func (s IngredientServer) DeleteIngredient(ctx context.Context, req *ingredientpb.DeleteIngredientRequest) (*ingredientpb.DeleteIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIngredient not implemented")
}
