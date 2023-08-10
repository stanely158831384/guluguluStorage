-- name: CreatePicture :one
INSERT INTO picture (
    link,
    user_id
) VALUES (
    $1, $2
) 
RETURNING *;

-- name: GetPicture :one
SELECT * FROM picture
WHERE id = $1 LIMIT 1;