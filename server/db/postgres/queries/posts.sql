-- name: ListPosts :many
SELECT * FROM posts;

-- name: CreatePost :one
INSERT INTO posts (
    user_id, content
) VALUES ($1, $2) RETURNING *;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: GetPopularPosts :many
SELECT *
FROM posts
ORDER BY views DESC
LIMIT $1 OFFSET $2;
