package main

import (
	"context"
	"fmt"

	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipeServer struct {
	recipepb.UnimplementedRecipeServiceServer
	db *sqlx.DB
}

func (s RecipeServer) CreateRecipe(ctx context.Context, req *recipepb.CreateRecipeRequest) (*recipepb.CreateRecipeResponse, error) {
	fmt.Println(req)
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
