-- name: GetUserById :one
SELECT id, name, email, password_hash, salt, created_at, verified_at, updated_at 
FROM users
WHERE id = $1;

-- name: InsertUser :exec
INSERT INTO users (id, name, email, password_hash, salt)
VALUES ($1, $2, $3, $4, $5);
