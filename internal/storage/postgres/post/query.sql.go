// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package post

import (
	"context"

	"github.com/lib/pq"
)

const createPost = `-- name: CreatePost :one
INSERT INTO post (caption, images)
VALUES ($1, $2)
RETURNING ident, caption, images, created_at
`

type CreatePostParams struct {
	Caption string
	Images  []string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.Caption, pq.Array(arg.Images))
	var i Post
	err := row.Scan(
		&i.Ident,
		&i.Caption,
		pq.Array(&i.Images),
		&i.CreatedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE
FROM post
WHERE ident = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, ident int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, ident)
	return err
}

const getPost = `-- name: GetPost :one
SELECT ident, caption, images, created_at
FROM post
WHERE ident = $1
LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, ident int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, ident)
	var i Post
	err := row.Scan(
		&i.Ident,
		&i.Caption,
		pq.Array(&i.Images),
		&i.CreatedAt,
	)
	return i, err
}

const getUserPosts = `-- name: GetUserPosts :many
SELECT ident, caption, images, created_at
FROM post
ORDER BY created_at
`

func (q *Queries) GetUserPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getUserPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.Ident,
			&i.Caption,
			pq.Array(&i.Images),
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}