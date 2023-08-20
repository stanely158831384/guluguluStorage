package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) (Product, error) {
	category, err := testStore.CreateCategory(context.Background(), "test")
	Ingredient, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: []string{"a", "b", "c", "d", "e", "f"},
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)
	require.NoError(t, err)
	require.NotEmpty(t, Ingredient)

	arg := CreateProductParams{
		Name: util.RandomString(6),
		CategoryID: category.ID,
		IngredientsID: Ingredient.ID,
	}
	product, err := testStore.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product, err
}


func TestCreateProduct(t *testing.T) {
	category, err := testStore.CreateCategory(context.Background(), "test")
	Ingredient, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: []string{"a", "b", "c", "d", "e", "f"},
		PictureID: pgtype.Int8{Int64: 1, Valid: true}},
	)
	arg := CreateProductParams{
		Name: util.RandomString(6),
		CategoryID: category.ID,
		IngredientsID: Ingredient.ID,
	}
	product, err := testStore.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)
}

func TestGetProduct(t *testing.T) {
	product1, err := createRandomProduct(t)
	require.NoError(t, err)
	require.NotEmpty(t, product1)

	product2, err := testStore.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Name, product2.Name)
}

func TestUpdateProduct(t *testing.T) {
	product1, err := createRandomProduct(t)

	require.NoError(t, err)
	require.NotEmpty(t, product1)

	category, err := testStore.CreateCategory(context.Background(), "category")
	ingredient, err := testStore.CreateIngredients(context.Background(), CreateIngredientsParams{
		Ingredient: []string{"a", "b", "c", "d", "e", "f"},
		PictureID: pgtype.Int8{Int64: 2, Valid: true}},
	)


	arg := UpdateProductParams{
		ID: product1.ID,
		Name: product1.Name,
		CategoryID: category.ID,
		IngredientsID: ingredient.ID,
	}
	product2, err := testStore.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, arg.Name, product2.Name)
	require.Equal(t, arg.CategoryID, product2.CategoryID)
	require.Equal(t, arg.IngredientsID, product2.IngredientsID)
}