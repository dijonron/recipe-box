package ingredient_service

import (
	"context"

	ingredientpb "github.com/dijonron/recipe-box/generated/ingredient_service/v1"
)

func (s IngredientServer) CreateIngredient(context.Context, *ingredientpb.CreateIngredientRequest) (*ingredientpb.IngredientResponse, error) {

	return &ingredientpb.IngredientResponse{
		Ingredient: &ingredientpb.Ingredient{
			Id:   "1",
			Name: "Water",
		},
	}, nil
}

func (s IngredientServer) GetIngredient(context.Context, *ingredientpb.GetIngredientRequest) (*ingredientpb.IngredientResponse, error) {

	return &ingredientpb.IngredientResponse{
		Ingredient: &ingredientpb.Ingredient{
			Id:   "1",
			Name: "Water",
		},
	}, nil
}
