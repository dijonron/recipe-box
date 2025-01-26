package main

import (
	"context"

	ingpb "github.com/dijonron/recipe-box/generated/proto/ingredient_service/v1"
)

func (s IngredientServer) CreateIngredient(context.Context, *ingpb.CreateIngredientRequest) (*ingpb.CreateIngredientResponse, error) {

	
	return &ingpb.CreateIngredientResponse{
		Ingredient: &ingpb.Ingredient{ 
			Id: "1",
			Name: "Water",
			Amount: 250,
			Measurement: ingpb.Measurement_MILLILITER,
		},
		
	}, nil
}

func (s IngredientServer) GetIngredient(context.Context, *ingpb.GetIngredientRequest) (*ingpb.GetIngredientResponse, error) {
	
	return &ingpb.GetIngredientResponse{
		Ingredient: &ingpb.Ingredient{
			Id: "1",
			Name: "Water",
			Amount: 250,
			Measurement: ingpb.Measurement_MILLILITER,
		},
		
	}, nil
}