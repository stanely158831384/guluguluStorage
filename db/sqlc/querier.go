// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, name string) (Category, error)
	CreateFeeling(ctx context.Context, arg CreateFeelingParams) (Feeling, error)
	CreateIngredients(ctx context.Context, arg CreateIngredientsParams) (Ingredient, error)
	CreatePicture(ctx context.Context, arg CreatePictureParams) (Picture, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	GetFeeling(ctx context.Context, id int64) (Feeling, error)
	GetIngredient(ctx context.Context, id int64) (Ingredient, error)
	GetPicture(ctx context.Context, id int64) (Picture, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
}

var _ Querier = (*Queries)(nil)