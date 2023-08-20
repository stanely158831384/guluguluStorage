package db

import (
	"context"
	"testing"

	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCategory(t *testing.T) {
	category, err := testStore.CreateCategory(context.Background(), "test")

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, "test", category.Name)
	require.NotZero(t, category.ID)
}

func TestGetCategory(t *testing.T) {
	category1, err := testStore.CreateCategory(context.Background(), "test")
	require.NoError(t, err)
	require.NotEmpty(t, category1)

	category2, err := testStore.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
}

func TestListCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, err := testStore.CreateCategory(context.Background(), "test")
		require.NoError(t, err)
	}
	categories, err := testStore.ListCategories(context.Background())
	require.NoError(t, err)


	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestDeleteCategory(t *testing.T) {
	category1, err := testStore.CreateCategory(context.Background(), "test")
	require.NoError(t, err)
	require.NotEmpty(t, category1)

	err = testStore.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testStore.GetCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
}

func TestUpdateCategory(t *testing.T) {
	str1 := util.RandomString(5)
	category1, err := testStore.CreateCategory(context.Background(), str1)
	require.NoError(t, err)
	require.NotEmpty(t, category1)

	str2 := util.RandomString(5)
	arg := UpdateCategoryParams{
		ID:   category1.ID,
		Name: str2,
	}

	category2, err := testStore.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, str2, category2.Name)
}

func createRandomCategory (t *testing.T) Category {
	str := util.RandomString(5)
	category, err := testStore.CreateCategory(context.Background(), str)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	return category
}


