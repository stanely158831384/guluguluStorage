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