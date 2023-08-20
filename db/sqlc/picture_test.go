package db

import (
	"context"
	"testing"

	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePicture (t *testing.T) {
	arg := CreatePictureParams{
		Link: util.RandomString(6),
		UserID: util.RandomAccountID(),
	}
	picture, err := testStore.CreatePicture(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, picture)

	require.Equal(t, arg.Link, picture.Link)
	require.Equal(t, arg.UserID, picture.UserID)
	require.NotZero(t, picture.ID)
	require.NotZero(t, picture.CreatedAt)
}

func TestGetPicture(t *testing.T) {
	arg := CreatePictureParams{
		Link: util.RandomString(6),
		UserID: util.RandomAccountID(),
	}
	picture1, err := testStore.CreatePicture(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, picture1)

	picture2, err := testStore.GetPicture(context.Background(), picture1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, picture2)

	require.Equal(t, picture1.ID, picture2.ID)
	require.Equal(t, picture1.Link, picture2.Link)
}

func TestDeletePicture(t *testing.T) {
	picture1, err := testStore.CreatePicture(context.Background(), CreatePictureParams{
		Link: util.RandomString(6),
		UserID: util.RandomAccountID(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, picture1)

	err = testStore.DeletePicture(context.Background(), picture1.ID)
	require.NoError(t, err)

	picture2, err := testStore.GetPicture(context.Background(), picture1.ID)
	require.Error(t, err)
	require.Empty(t, picture2)
}

