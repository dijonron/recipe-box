package main

import (
	"context"

	"github.com/dijonron/recipe-box/generated/proto/ingredient_service/v1/ingredientpb"
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
