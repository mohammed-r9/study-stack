-- name: InsertToken :exec
INSERT INTO tokens (hash, user_id, scope, expires_at)
VALUES( $1, $2, $3, $4);

-- name: GetTokenByHash :one
SELECT hash, user_id, scope, created_at, used_at, revoked_at, expires_at
FROM tokens
WHERE hash = $1;

-- name: UseToken :execrows
UPDATE tokens
SET used_at = CURRENT_TIMESTAMP
WHERE hash = $1;

-- name: GetAllTokensByUserID :many
SELECT hash, user_id, scope, created_at, used_at, revoked_at, expires_at
FROM tokens
WHERE user_id = $1;

-- name: GetAllTokensByScope :many
SELECT
    tokens.hash,
    tokens.user_id,
    tokens.scope,
    tokens.created_at,
    tokens.used_at,
    tokens.revoked_at,
    tokens.expires_at
FROM tokens
JOIN users
    ON users.id = tokens.user_id
WHERE users.email = $1
    AND tokens.scope = $2
    AND tokens.revoked_at IS NULL
    AND tokens.expires_at > NOW()
ORDER BY tokens.used_at DESC NULLS LAST;
