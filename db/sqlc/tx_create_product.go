package db

import "context"

type CreateProductTxParams struct {
	CreateProductParams
	AfterCreate func(product Product) error
}

type CreateProductTxResult struct {
	Product Product
}

func (store *SQLStore) CreateProductTx(ctx context.Context, arg CreateProductTxParams) (CreateProductTxResult, error){
	var result CreateProductTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Product, err = q.CreateProduct(ctx, arg.CreateProductParams)

		if err != nil {
			return err
		}


		err = arg.AfterCreate(result.Product)
		return err
	})

	return result, err
}