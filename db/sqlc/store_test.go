package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)

func createRandomPicture(t *testing.T) Picture {
	arg := CreatePictureParams{
		Link: util.RandomString(6),
		Username: util.RandomString(6),
	}
	picture, err := testStore.CreatePicture(context.Background(), arg)
	require.NoError(t, err)
	return picture
}

func test1(t *testing.T,errs chan error,resultsIngredients chan CreateIngredientsTxResult,picture Picture,ingredients []string,category Category){

	ctx1, cancel1 := context.WithCancel(context.Background())
	idx1 := 0
	for {
		select {
			case <-ctx1.Done():
				return
			default:
				CreateIngredientsParams := CreateIngredientParams{
					PictureID: pgtype.Int8{picture.ID, true},
					Ingredient: ingredients,
				}
				result, err := testStore.CreateIngredientsTx(context.Background(), CreateIngredientsTxParams{
						CreateIngredientParams: CreateIngredientsParams,
						AfterCreate: func(Ingredient Ingredient) error {
							require.NotEmpty(t, Ingredient)
							return nil
						},
				})
				require.NoError(t, err)
				errs <- err
				resultsIngredients <- result
				idx1++
				if idx1 >= 5 {
					cancel1()
				}
		}
	}

}

func test2(t *testing.T, ctx context.Context, errs chan error, resultsIngredients chan CreateIngredientsTxResult, resultsProducts chan CreateProductTxResult,category Category) {
	ctx2, cancel2 := context.WithCancel(context.Background())
	idx2 := 0
	for {
		select {
			case <-ctx2.Done():
				return
			default:
				err := <-errs
				resultForIngredient := <-resultsIngredients
				CreateProductParams := CreateProductParams{
					Name: util.RandomString(6),
					CategoryID: category.ID,
					IngredientsID: resultForIngredient.Ingredient.ID,
				}
				result, err := testStore.CreateProductTx(context.Background(), CreateProductTxParams{
					CreateProductParams: CreateProductParams,
					AfterCreate: func(Product Product) error {
						require.NotEmpty(t, Product)
						return nil
					},
				})
				require.NoError(t, err)
				resultsProducts<-result
				idx2++
				if idx2 >= 5 {
					cancel2()
				}
		}
	}
}
func TestCreateProductTx3(t *testing.T) {
	errs := make(chan error)
	resultsIngredients := make(chan CreateIngredientsTxResult)
	resultsProducts := make(chan CreateProductTxResult)
	picture := createRandomPicture(t)
	ingredients := []string{"a", "b", "c", "d", "e"}
	category := createRandomCategory(t)


	go test1(t,errs,resultsIngredients,picture,ingredients,category)
	go test2(t,context.Background(),errs,resultsIngredients,resultsProducts,category)
	for i:=0; i<5; i++{
		result := <-resultsProducts
		require.NotEmpty(t, result.Product)
	}
}



func TestCreateProductTx2(t *testing.T) {
	errs := make(chan error)
	resultsIngredients := make(chan CreateIngredientsTxResult)
	resultsProducts := make(chan CreateProductTxResult)
	picture := createRandomPicture(t)
	ingredients := []string{"a", "b", "c", "d", "e"}
	category := createRandomCategory(t)

	for i:=0; i<5; i++{
		CreateIngredientsParams := CreateIngredientParams{
			PictureID: pgtype.Int8{picture.ID, true},
			Ingredient: ingredients,
		}
		go func() {
			result, err := testStore.CreateIngredientsTx(context.Background(), CreateIngredientsTxParams{
				CreateIngredientParams: CreateIngredientsParams,
				AfterCreate: func(Ingredient Ingredient) error {
					require.NotEmpty(t, Ingredient)
					return nil
				},
			})
			require.NoError(t, err)
			errs <- err
			resultsIngredients <- result
		}()
	}

	for i:=0; i<5; i++{



		go func() {
			err := <-errs
			resultForIngredient := <-resultsIngredients
			CreateProductParams := CreateProductParams{
				Name: util.RandomString(6),
				CategoryID: category.ID,
				IngredientsID: resultForIngredient.Ingredient.ID,
			}
			result, err := testStore.CreateProductTx(context.Background(), CreateProductTxParams{
				CreateProductParams: CreateProductParams,
				AfterCreate: func(Product Product) error {
					require.NotEmpty(t, Product)
					return nil
				},
			})
			require.NoError(t, err)
			resultsProducts<-result
		}()
	}

	for i:=0; i<5; i++{
		result := <-resultsProducts
		require.NotEmpty(t, result.Product)
	}
	

}

// func TestCreateProductsTx(t *testing.T) {
// 	picture1 := createRandomPicture(t)
// 	n := 5
// 	errs := make(chan error)
// 	resultsIngredients := make(chan CreateIngredientsTxResult)
// 	resultsProducts := make(chan CreateProductTxResult)

// 	ingredients := []string{"a", "b", "c", "d", "e"}
// 	for i := 0; i < n; i++ {
// 		CreateIngredientsParams := CreateIngredientParams{
// 			PictureID: pgtype.Int8{picture1.ID, true},
// 			Ingredient: ingredients,
// 		}
// 		go func() {
// 			result, err := testStore.CreateIngredientsTx(context.Background(), CreateIngredientsTxParams{
// 				CreateIngredientParams: CreateIngredientsParams,
// 				AfterCreate: func(Ingredient Ingredient) error {
// 					require.NotEmpty(t, Ingredient)
// 					return nil
// 				},
// 			})
// 			errs <- err
// 			resultsIngredients <- result
// 		}()
// 	}



// 	category := createRandomCategory(t)
// 	categoryID := category.ID

// 	for i := 0; i < n; i++ {
// 		result := <-resultsIngredients
// 		CreateProductParams := CreateProductParams{
// 			Name: util.RandomString(6),
// 			CategoryID: category.ID,
// 			IngredientsID: result.Ingredient.ID,
// 		}

// 		go func() {

// 			result, err := testStore.CreateProductTx(context.Background(), CreateProductTxParams{
// 				CreateProductParams: CreateProductParams,
// 				AfterCreate: func(Product Product) error {
// 					require.NotEmpty(t, Product)
// 					return nil
// 				},
// 			})
// 			errs <- err
// 			resultsProducts <- result
// 		}()
// 	}


// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)
// 		result := <-resultsProducts
// 		require.Equal(t, result.Product.CategoryID, categoryID)
// 	}


// }


