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

-- name: DeletePost :exec
DELETE
FROM post
WHERE ident = $1;

-- name: ListPosts :many
SELECT *
FROM post
order by created_at;

-- name: GetPostAndUser :one
select p.*, "user".*
from (select * from "post" where ident = $1) p
         join "user" on "p".owner = "user".ident
limit 1;