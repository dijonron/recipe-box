package main

import (
	"context"

	ingredientpb "github.com/dijonron/recipe-box/server/generated/ingredient_service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IngredientServer struct {
	ingredientpb.UnimplementedIngredientServiceServer
}

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

func (s IngredientServer) ListIngredients(context.Context, *ingredientpb.ListIngredientsRequest) (*ingredientpb.ListIngredientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIngredients not implemented")
}
func (s IngredientServer) UpdateIngredient(context.Context, *ingredientpb.UpdateIngredientRequest) (*ingredientpb.IngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}
func (s IngredientServer) DeleteIngredient(context.Context, *ingredientpb.DeleteIngredientRequest) (*ingredientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIngredient not implemented")
}
