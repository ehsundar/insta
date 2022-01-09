-- name: GetPost :one
SELECT *
FROM post
WHERE ident = $1
LIMIT 1;

-- name: GetUserPosts :many
SELECT *
FROM post
ORDER BY created_at;

-- name: CreatePost :one
INSERT INTO post (caption, images)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE
FROM post
WHERE ident = $1;