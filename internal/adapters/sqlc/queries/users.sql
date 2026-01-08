-- name: GetUserByID :one
SELECT id, name, email, password_hash, salt, created_at, verified_at, updated_at 
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, name, email, password_hash, salt, created_at, verified_at, updated_at 
FROM users
WHERE email = $1;

-- name: InsertUser :exec
INSERT INTO users (id, name, email, password_hash, salt)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateUserName :exec
UPDATE users
SET name = $1
WHERE id = $2;

-- name: UpdateUserEmail :exec
UPDATE users
SET email = $1
WHERE id = $2;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $1, salt = $2
WHERE id = $3;

-- name: verifyUserEmail :exec
UPDATE users
SET verified_at = NOW()
WHERE id = $1;
