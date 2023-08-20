package db

import "context"

type CreateIngredientsTxParams struct {
	CreateIngredientsParams
	AfterCreate func(Ingredient Ingredient) error
}

type CreateIngredientsTxResult struct {
	Ingredient Ingredient
}

func (store *SQLStore) CreateIngredientsTx(ctx context.Context, arg CreateIngredientsTxParams) (CreateIngredientsTxResult, error){
	var result CreateIngredientsTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Ingredient, err = q.CreateIngredients(ctx, arg.CreateIngredientsParams)

		if err != nil {
			return err
		}


		err = arg.AfterCreate(result.Ingredient)
		return err
	})

	return result, err
}