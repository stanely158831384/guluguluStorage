// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: product.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product (
    name,
    category_id,
    ingredients_id
) VALUES (
    $1, $2, $3
) 
RETURNING id, name, category_id, ingredients_id, risk_level, picture_id, created_at
`

type CreateProductParams struct {
	Name          string `json:"name"`
	CategoryID    int64  `json:"category_id"`
	IngredientsID int64  `json:"ingredients_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.Name, arg.CategoryID, arg.IngredientsID)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.IngredientsID,
		&i.RiskLevel,
		&i.PictureID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, category_id, ingredients_id, risk_level, picture_id, created_at FROM product
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.IngredientsID,
		&i.RiskLevel,
		&i.PictureID,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, category_id, ingredients_id, risk_level, picture_id, created_at FROM product
ORDER BY id
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CategoryID,
			&i.IngredientsID,
			&i.RiskLevel,
			&i.PictureID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE product SET name = $2, category_id = $3, ingredients_id = $4
WHERE id = $1
RETURNING id, name, category_id, ingredients_id, risk_level, picture_id, created_at
`

type UpdateProductParams struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	CategoryID    int64  `json:"category_id"`
	IngredientsID int64  `json:"ingredients_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.CategoryID,
		arg.IngredientsID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.IngredientsID,
		&i.RiskLevel,
		&i.PictureID,
		&i.CreatedAt,
	)
	return i, err
}
