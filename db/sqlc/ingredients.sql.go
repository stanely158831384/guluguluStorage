// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: ingredients.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createIngredients = `-- name: CreateIngredients :one
INSERT INTO ingredients (
    ingredient,
    picture_id
) VALUES (
    $1, $2
) 
RETURNING id, ingredient, picture_id, created_at
`

type CreateIngredientsParams struct {
	Ingredient []string    `json:"ingredient"`
	PictureID  pgtype.Int8 `json:"picture_id"`
}

func (q *Queries) CreateIngredients(ctx context.Context, arg CreateIngredientsParams) (Ingredient, error) {
	row := q.db.QueryRow(ctx, createIngredients, arg.Ingredient, arg.PictureID)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.Ingredient,
		&i.PictureID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteIngredient = `-- name: DeleteIngredient :exec
DELETE FROM ingredients WHERE id = $1
`

func (q *Queries) DeleteIngredient(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteIngredient, id)
	return err
}

const getIngredient = `-- name: GetIngredient :one
SELECT id, ingredient, picture_id, created_at FROM ingredients
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetIngredient(ctx context.Context, id int64) (Ingredient, error) {
	row := q.db.QueryRow(ctx, getIngredient, id)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.Ingredient,
		&i.PictureID,
		&i.CreatedAt,
	)
	return i, err
}

const listIngredients = `-- name: ListIngredients :many
SELECT id, ingredient, picture_id, created_at FROM ingredients
ORDER BY id
`

func (q *Queries) ListIngredients(ctx context.Context) ([]Ingredient, error) {
	rows, err := q.db.Query(ctx, listIngredients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ingredient{}
	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(
			&i.ID,
			&i.Ingredient,
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

const updateIngredient = `-- name: UpdateIngredient :one
UPDATE ingredients SET ingredient = $2, picture_id = $3
WHERE id = $1
RETURNING id, ingredient, picture_id, created_at
`

type UpdateIngredientParams struct {
	ID         int64       `json:"id"`
	Ingredient []string    `json:"ingredient"`
	PictureID  pgtype.Int8 `json:"picture_id"`
}

func (q *Queries) UpdateIngredient(ctx context.Context, arg UpdateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRow(ctx, updateIngredient, arg.ID, arg.Ingredient, arg.PictureID)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.Ingredient,
		&i.PictureID,
		&i.CreatedAt,
	)
	return i, err
}
