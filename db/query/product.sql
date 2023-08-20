-- name: CreateProduct :one
INSERT INTO product (
    name,
    category_id,
    ingredients_id
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM product
ORDER BY id;

-- name: UpdateProduct :one
UPDATE product SET name = $2, category_id = $3, ingredients_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;