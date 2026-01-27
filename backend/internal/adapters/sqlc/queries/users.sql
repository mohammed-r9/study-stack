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

-- name: UpdateUserName :execrows
UPDATE users
SET name = $1
WHERE id = $2;

-- name: UpdateUserEmail :execrows
UPDATE users
SET email = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserPassword :execrows
UPDATE users
SET password_hash = $1, salt = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $3;

-- name: verifyUserEmail :execrows
UPDATE users
SET verified_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
