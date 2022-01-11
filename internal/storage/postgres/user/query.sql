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
