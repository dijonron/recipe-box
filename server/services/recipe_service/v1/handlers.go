package main

import (
	"context"

	"github.com/dijonron/recipe-box/generated/proto/recipe_service/v1/recipepb"
)

func (s RecipeServer) CreateRecipe(context.Context, *recipepb.CreateRecipeRequest) (*recipepb.RecipeResponse, error) {

	return &recipepb.RecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          "1",
			Name:        "Cookies",
			Chef:        "Me",
			Cookbook:    "Yes",
			Ingredients: []*recipepb.RecipeIngredient{{IngredientId: "1", Name: "Flour", Amount: 3, Measurement: recipepb.Measurement_CUP}},
		},
	}, nil
}

func (s RecipeServer) GetRecipe(context.Context, *recipepb.GetRecipeRequest) (*recipepb.RecipeResponse, error) {

	return &recipepb.RecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          "1",
			Name:        "Cookies",
			Chef:        "Me",
			Cookbook:    "Yes",
			Ingredients: []*recipepb.RecipeIngredient{{IngredientId: "1", Name: "Flour", Amount: 3, Measurement: recipepb.Measurement_CUP}},
		},
	}, nil
}
