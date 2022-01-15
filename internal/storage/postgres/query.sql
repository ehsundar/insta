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
from (select * from "post" where post.ident = $1) p
         join "user" on "p".owner = "user".ident
limit 1;

-- name: CreateUser :one
insert into "user"
(username, password)
values ($1, $2)
returning ident;

-- name: GetUserByIdent :one
select *
from "user"
where ident = $1
limit 1;

-- name: GetUserByUsername :one
select *
from "user"
where username = $1
limit 1;
