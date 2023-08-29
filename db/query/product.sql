-- name: CreateProduct :one
INSERT INTO product (
    name,
    category_id,
    ingredients_id,
    username    
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM product
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListProductsByCategory :many
SELECT * FROM product
WHERE category_id = $1
LIMIT $2
OFFSET $3;

-- name: ListProductsByUserID :many
SELECT * FROM product
WHERE username = $1
LIMIT $2
OFFSET $3;

-- name: UpdateProduct :one
UPDATE product SET name = $2, category_id = $3, ingredients_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;