// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (
    name
) VALUES (
    $1
) 
RETURNING id, name
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name FROM category
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRow(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
