-- name: NewUserSession :exec
INSERT INTO sessions(id, user_id, token_hash, device_name, csrf_hash)
VALUES($1, $2, $3, $4, $5);

-- name: GetSessionByHash :one
SELECT sessions.id, 
    sessions.user_id, 
    sessions.token_hash, 
    sessions.csrf_hash, 
    sessions.device_name, 
    sessions.last_used_at, 
    sessions.revoked_at,
    users.verified_at
FROM sessions
JOIN users ON users.id = sessions.user_id
WHERE token_hash = $1;
