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

-- name: ListPictures :many
SELECT * FROM picture
ORDER BY id;

-- name: UpdatePicture :one
UPDATE picture SET link = $2, user_id = $3
WHERE id = $1
RETURNING *;

-- name: DeletePicture :exec
DELETE FROM picture WHERE id = $1;