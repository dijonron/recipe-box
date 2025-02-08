package recipe_persistence

import (
	"fmt"
	"log"

	recipepb "github.com/dijonron/recipe-box/server/generated/recipe_service/v1"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func WriteRecipe(recipe Recipe, db *sqlx.DB) (*recipepb.Recipe, error) {
	var newRecipe []recipepb.Recipe

	tx, err := db.Beginx()
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

	query := `
		INSERT INTO rs_recipes (name, chef, cookbook)
		VALUES ($1, $2, $3)
		RETURNING id, name, chef, cookbook
	`

	err = tx.Select(&newRecipe, query, recipe.Name, recipe.Chef, recipe.Cookbook)
	if err != nil {
		log.Println("Error inserting recipe:", err)
		return nil, fmt.Errorf("failed to insert ingredients: %w", err)
	}

	return &newRecipe[0], nil
}

func WriteRecipeIngredients(ingredients []RecipeIngredient, db *sqlx.DB) error {
	var insertedIngredients []recipepb.RecipeIngredient

	tx, err := db.Beginx()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := `
		INSERT INTO rs_recipe_ingredients (ingredient_id, recipe_id, amount, measurement)
		SELECT * FROM unnest($1::UUID[], $2::UUID[], $3::float8[], $4::integer[]) AS t(ingredient_id, recipe_id, amount, measurement)
	`

	numIng := len(ingredients)
	ingredientIds := make([]uuid.UUID, 0, numIng)
	recipeIds := make([]uuid.UUID, 0, numIng)
	amounts := make([]float64, 0, numIng)
	measurement := make([]float64, 0, numIng)

	for _, ing := range ingredients {
		ingredientIds = append(ingredientIds, ing.IngredientId)
		recipeIds = append(recipeIds, ing.RecipeID)
		amounts = append(amounts, ing.Amount)
		measurement = append(measurement, float64(ing.Measurement))
	}

	err = tx.Select(&insertedIngredients, query, pq.Array(ingredientIds), pq.Array(recipeIds), pq.Array(amounts), pq.Array(measurement))
	if err != nil {
		log.Println("Error inserting ingredients:", err)
		return fmt.Errorf("failed to insert ingredients: %w", err)
	}

	return nil
}

func DeleteRecipe(recipeId string, tx *sqlx.Tx) error {
	query := `
		UPDATE rs_recipes
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1	
	`

	_, err := tx.Exec(query, recipeId)
	if err != nil {
		log.Println("Error deleting recipe:", err)
		return fmt.Errorf("failed to delete recipe: %w", err)
	}

	return nil
}

func DeleteRecipeIngredients(recipeId string, tx *sqlx.Tx) error {
	query := `
		UPDATE rs_recipe_ingredients 
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE recipe_id = $1	
	`

	_, err := tx.Exec(query, recipeId)
	if err != nil {
		log.Println("Error deleting ingredients:", err)
		return fmt.Errorf("failed to delete ingredients: %w", err)
	}

	return nil
}
