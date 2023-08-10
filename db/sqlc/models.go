// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Feeling struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Comment   string    `json:"comment"`
	Recommend bool      `json:"recommend"`
	CreatedAt time.Time `json:"created_at"`
}

type Ingredient struct {
	ID         int64       `json:"id"`
	Ingredient []string    `json:"ingredient"`
	PictureID  pgtype.Int8 `json:"picture_id"`
	CreatedAt  time.Time   `json:"created_at"`
}

type Picture struct {
	ID        int64       `json:"id"`
	Link      pgtype.Text `json:"link"`
	UserID    pgtype.Int8 `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
}

type Product struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	CategoryID    int64       `json:"category_id"`
	IngredientsID int64       `json:"ingredients_id"`
	RiskLevel     pgtype.Int2 `json:"risk_level"`
	PictureID     pgtype.Int8 `json:"picture_id"`
	CreatedAt     time.Time   `json:"created_at"`
}
