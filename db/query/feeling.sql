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

-- name: ListFeelings :many
SELECT * FROM feeling
ORDER BY id;

-- name: UpdateFeeling :one
UPDATE feeling SET product_id = $2, user_id = $3, username = $4, comment = $5, recommend = $6
WHERE id = $1
RETURNING *;

-- name: DeleteFeeling :exec
DELETE FROM feeling WHERE id = $1;
