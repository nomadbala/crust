-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT  1;

-- name: GetUser :one
SELECT id, password_hash, salt FROM users
WHERE username = $1;

-- name: GetEmailById :one
SELECT users.email FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
    username, password_hash, salt, email
) VALUES ($1, $2, $3, $3) RETURNING *;