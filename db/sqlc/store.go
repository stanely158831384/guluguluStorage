package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateProductTx(ctx context.Context, arg CreateProductTxParams) (CreateProductTxResult, error)
	CreateIngredientsTx(ctx context.Context, arg CreateIngredientsTxParams) (CreateIngredientsTxResult, error)
}

// SQLSTORE provides all functions to execute sql queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

//New Store creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries: New(connPool),
	}
}

