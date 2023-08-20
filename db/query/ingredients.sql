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

-- name: ListIngredients :many
SELECT * FROM ingredients
ORDER BY id;

-- name: UpdateIngredient :one
UPDATE ingredients SET ingredient = $2, picture_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM ingredients WHERE id = $1;