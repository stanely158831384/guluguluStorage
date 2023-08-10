-- name: CreateIngredients :one
INSERT INTO ingredients (
    ingredient,
    picture_id
) VALUES (
    $1, $2
) 
RETURNING *;

-- name: GetIngredient :one
SELECT * FROM ingredients
WHERE id = $1 LIMIT 1;