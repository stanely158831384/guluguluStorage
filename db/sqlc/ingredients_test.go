package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateIngredients(t *testing.T) {
	arg := []string{"a", "b", "c", "d", "e", "f"};
	ingredient, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: arg,
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient)

	require.Equal(t, arg[0], ingredient.Ingredient[0])
	require.Equal(t, arg[1], ingredient.Ingredient[1])
	require.NotZero(t, ingredient.PictureID)
}

func TestGetIngredients(t *testing.T) {
	arg := []string{"a", "b", "c", "d", "e", "f"};
	ingredient1, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: arg,
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)
	require.NoError(t, err)
	require.NotEmpty(t, ingredient1)

	ingredient2, err := testStore.GetIngredient(context.Background(), ingredient1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, ingredient2)

	require.Equal(t, ingredient1.ID, ingredient2.ID)
	require.Equal(t, ingredient1.Ingredient[0], ingredient2.Ingredient[0])
}

func TestDeleteIngredients(t *testing.T) {
	arg := []string{"a", "b", "c", "d", "e", "f"};
	ingredient1, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: arg,
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)
	require.NoError(t, err)
	require.NotEmpty(t, ingredient1)

	err = testStore.DeleteIngredient(context.Background(), ingredient1.ID)
	require.NoError(t, err)

	ingredient2, err := testStore.GetIngredient(context.Background(), ingredient1.ID)
	require.Error(t, err)
	require.Empty(t, ingredient2)
}

func TestUpdateIngredient(t *testing.T) {
	arg := []string{"a", "b", "c", "d", "e", "f"};
	ingredient1, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: arg,
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)
	require.NoError(t, err)
	require.NotEmpty(t, ingredient1)

	arg2 := []string{"a", "b", "c", "d", "e", "f"};
	ingredient2, err := testStore.UpdateIngredient(context.Background(), UpdateIngredientParams{
		ID: ingredient1.ID,
		Ingredient: arg2,
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)
	require.NoError(t, err)
	require.NotEmpty(t, ingredient2)

	require.Equal(t, ingredient1.ID, ingredient2.ID)
	require.Equal(t, ingredient1.Ingredient[0], ingredient2.Ingredient[0])
	require.Equal(t, ingredient1.Ingredient[1], ingredient2.Ingredient[1])
	require.Equal(t, ingredient1.Ingredient[2], ingredient2.Ingredient[2])
	require.Equal(t, ingredient1.Ingredient[3], ingredient2.Ingredient[3])
	require.Equal(t, ingredient1.Ingredient[4], ingredient2.Ingredient[4])
	require.Equal(t, ingredient1.Ingredient[5], ingredient2.Ingredient[5])
}