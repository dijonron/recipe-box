package main

import (
	"context"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
	models "github.com/dijonron/recipe-box/server/services/ingredient_service/v1/persistence"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IngredientServer struct {
	ingredientpb.UnimplementedIngredientServiceServer
	db *sqlx.DB
}

func (s IngredientServer) CreateIngredient(ctx context.Context, req *ingredientpb.CreateIngredientRequest) (*ingredientpb.CreateIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}

func (s IngredientServer) GetIngredient(ctx context.Context, req *ingredientpb.GetIngredientRequest) (*ingredientpb.GetIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}

func (s IngredientServer) ListIngredients(ctx context.Context, req *ingredientpb.ListIngredientsRequest) (*ingredientpb.ListIngredientsResponse, error) {
	var ingredients []models.Ingredient

	err := s.db.Select(&ingredients, "SELECT id, name FROM is_ingredients WHERE deleted_at IS NULL ORDER BY name ASC")
	if err != nil {
		return nil, err
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

func (s IngredientServer) UpdateIngredient(ctx context.Context, req *ingredientpb.UpdateIngredientRequest) (*ingredientpb.UpdateIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}
func (s IngredientServer) DeleteIngredient(ctx context.Context, req *ingredientpb.DeleteIngredientRequest) (*ingredientpb.DeleteIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIngredient not implemented")
}
