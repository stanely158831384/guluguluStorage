package db

import (
	"context"
	"testing"

	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)

func TestCreateFeeling(t *testing.T){
	product, err := createRandomProduct(t)
	require.NoError(t, err)

	arg := CreateFeelingParams{
		ProductID: product.ID,
		UserID: 1,
		Username: "test",
		Comment: "this is a test",
		Recommend: true,
	}
	feeling, err := testStore.CreateFeeling(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, feeling)

	require.Equal(t, "test", feeling.Username)
	require.NotZero(t, feeling.ID)
}

func TestGetFeeling(t *testing.T) {
	product, err := createRandomProduct(t)
	require.NoError(t, err)
	feeling1, err := testStore.CreateFeeling(context.Background(), CreateFeelingParams{
		ProductID: product.ID,
		UserID: 1,
		Username: "test",
		Comment: "this is a test",
		Recommend: true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, feeling1)

	feeling2, err := testStore.GetFeeling(context.Background(), feeling1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, feeling2)

	require.Equal(t, feeling1.ID, feeling2.ID)
	require.Equal(t, feeling1.Username, feeling2.Username)
}

func TestListFeelings(t *testing.T) {
	productArray := [10]Product{}
	for i:=0; i<10; i++ {
		product, err := createRandomProduct(t)
		require.NoError(t, err)
		productArray[i] = product
	}
	for i := 0; i < 10; i++ {
		_, err := testStore.CreateFeeling(context.Background(), CreateFeelingParams{
			ProductID: productArray[i].ID,
			UserID: util.RandomAccountID(),
			Username: "test",
			Comment: "this is a test",
			Recommend: true,
		})
		require.NoError(t, err)
	}

	arg := ListFeelingsParams{
		Limit:  5,
		Offset: 0,
	}

	feelings, err := testStore.ListFeelings(context.Background(),arg)
	require.NoError(t, err)

	for _, feeling := range feelings {
		require.NotEmpty(t, feeling)
	}
}

func TestDeleteFeeling(t *testing.T) {
	product, err := createRandomProduct(t)
	require.NoError(t, err)
	feeling1, err := testStore.CreateFeeling(context.Background(), CreateFeelingParams{
		ProductID: product.ID,
		UserID: 1,
		Username: "test",
		Comment: "this is a test",
		Recommend: true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, feeling1)

	err = testStore.DeleteFeeling(context.Background(), feeling1.ID)
	require.NoError(t, err)

	feeling2, err := testStore.GetFeeling(context.Background(), feeling1.ID)
	require.Error(t, err)
	require.Empty(t, feeling2)
}

func TestUpdateFeeling(t *testing.T) {
	product, err := createRandomProduct(t)
	require.NoError(t, err)
	str1 := util.RandomString(5)
	feeling1, err := testStore.CreateFeeling(context.Background(), CreateFeelingParams{
		ProductID: product.ID,
		UserID: 1,
		Username: str1,
		Comment: "this is a test",
		Recommend: true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, feeling1)

	str2 := util.RandomString(5)
	arg := UpdateFeelingParams{
		ID:   feeling1.ID,
		Username: str2,
		ProductID: product.ID,
	}
	feeling2, err := testStore.UpdateFeeling(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, feeling2)

	require.Equal(t, feeling1.ID, feeling2.ID)
	require.Equal(t, str2, feeling2.Username)
}
