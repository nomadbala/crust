-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    username, password_hash, salt, email
) VALUES ($1, $2, $3, $3) RETURNING *;