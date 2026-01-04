-- name: GetUserById :one
SELECT id, name, email, password_hash, salt, created_at, verified_at, updated_at 
FROM users
WHERE id = $1;
