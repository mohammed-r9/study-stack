-- name: CreateCollection :execrows
INSERT INTO collections (id, user_id, title, description)
SELECT
    $1, $2, $3, $4
WHERE (
    SELECT COUNT(*)
    FROM collections
    WHERE user_id = $2
      AND archived_at IS NULL
) < 20;

-- name: ArchiveCollection :execrows
UPDATE collections
SET archived_at = NOW()
WHERE id = $1
  AND user_id = $2
  AND archived_at IS NULL;

-- name: UnarchiveCollection :execrows
UPDATE collections
SET archived_at = NULL
WHERE id = $1
  AND user_id = $2
  AND archived_at IS NOT NULL;

-- name: GetArchivedCollectionByID :one
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE id = $1
  AND user_id = $2
  AND archived_at IS NOT NULL;

-- name: GetArchivedCollections :many
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE user_id = $1
  AND archived_at IS NOT NULL
ORDER BY archived_at DESC
LIMIT $2 OFFSET $3;

-- name: GetCollectionByID :one
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE id = $1
  AND user_id = $2
  AND archived_at IS NULL;

-- name: GetAllCollections :many
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE user_id = $1
  AND archived_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
