-- name: CreateFeeling :one
INSERT INTO feeling (
    product_id,
    user_id,
    username,
    comment,
    recommend
) VALUES (
    $1, $2, $3, $4, $5
) 
RETURNING *;

-- name: GetFeeling :one
SELECT * FROM feeling
WHERE id = $1 LIMIT 1;