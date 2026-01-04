-- name: NewUserSession :exec
INSERT INTO sessions(id, user_id, token_hash, device_name, csrf_hash)
VALUES($1, $2, $3, $4, $5);

-- name: GetSessionByHash :one
SELECT id, user_id, token_hash, csrf_hash, device_name, last_used_at, revoked_at
FROM sessions
WHERE token_hash = $1;
