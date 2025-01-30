package main

import (
	"context"

	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipeServer struct {
	recipepb.UnimplementedRecipeServiceServer
}

func (s RecipeServer) CreateRecipe(context.Context, *recipepb.CreateRecipeRequest) (*recipepb.CreateRecipeResponse, error) {

	return &recipepb.CreateRecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          "1",
			Name:        "Cookies",
			Chef:        "Me",
			Cookbook:    "Yes",
			Ingredients: []*recipepb.RecipeIngredient{{IngredientId: "1", Name: "Flour", Amount: 3, Measurement: recipepb.Measurement_MEASUREMENT_CUP}},
		},
	}, nil
}

func (s RecipeServer) GetRecipe(context.Context, *recipepb.GetRecipeRequest) (*recipepb.GetRecipeResponse, error) {

	return &recipepb.GetRecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          "1",
			Name:        "Cookies",
			Chef:        "Me",
			Cookbook:    "Yes",
			Ingredients: []*recipepb.RecipeIngredient{{IngredientId: "1", Name: "Flour", Amount: 3, Measurement: recipepb.Measurement_MEASUREMENT_CUP}},
		},
	}, nil
}

func (s RecipeServer) ListRecipes(context.Context, *recipepb.ListRecipesRequest) (*recipepb.ListRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipes not implemented")
}
func (s RecipeServer) UpdateRecipe(context.Context, *recipepb.UpdateRecipeRequest) (*recipepb.UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
func (s RecipeServer) DeleteRecipe(context.Context, *recipepb.DeleteRecipeRequest) (*recipepb.DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
