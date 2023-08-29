-- name: CreateCategory :one
INSERT INTO category (
    name
) VALUES (
    $1
) 
RETURNING *;

-- name: GetCategory :one
SELECT * FROM category
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM category
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateCategory :one
UPDATE category SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;
