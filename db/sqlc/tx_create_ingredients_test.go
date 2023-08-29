package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)


func TestCreateIngredientsTx(t *testing.T) {
	// arg := []string{"a", "b", "c", "d", "e", "f"};
	// pgtype.Int8{Int64: 1, Valid: true}}
	arg := CreateIngredientsTxParams{
		CreateIngredientParams: CreateIngredientParams{
			Ingredient: []string{"a", "b", "c", "d", "e", "f"},
			PictureID: pgtype.Int8{Int64: 1, Valid: true}},
		AfterCreate: func(Ingredient Ingredient) error {
			require.NotEmpty(t, Ingredient)
			return nil
		},
	}

	result, err := testStore.CreateIngredientsTx(context.Background(), arg)

	require.Equal(t, arg.CreateIngredientParams.Ingredient, result.Ingredient.Ingredient)

	require.NoError(t, err)
	require.NotEmpty(t, result)
}


func TestMultiDeadlock(t *testing.T){
	go TestCreateIngredientsTx(t)
	go TestCreateIngredientsTx(t)
	go TestCreateIngredientsTx(t)
	go TestCreateIngredientsTx(t)
	go TestCreateIngredientsTx(t)

}