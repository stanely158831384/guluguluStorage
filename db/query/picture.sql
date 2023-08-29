-- name: CreatePicture :one
INSERT INTO picture (
    link,
    username
) VALUES (
    $1, $2
) 
RETURNING *;

-- name: GetPicture :one
SELECT * FROM picture
WHERE id = $1 LIMIT 1;

-- name: ListPictures :many
SELECT * FROM picture
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePicture :one
UPDATE picture SET link = $2, username = $3
WHERE id = $1
RETURNING *;

-- name: DeletePicture :exec
DELETE FROM picture WHERE id = $1;

-- name: ListPicturesByUsername :many
SELECT * FROM picture
WHERE username = $1
ORDER BY id
LIMIT $2
OFFSET $3;