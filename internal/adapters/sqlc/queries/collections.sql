-- name: CreateCollection :execrows
INSERT INTO collections(id, user_id, title, description)
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
WHERE collections.id = $1
  AND collections.user_id = $2
  AND collections.archived_at IS NOT NULL
  AND (
      SELECT COUNT(*)
      FROM collections AS active
      WHERE active.user_id = collections.user_id
        AND active.archived_at IS NULL
  ) < 20;

-- name: UpdateCollectionTitle :execrows
UPDATE collections
    SET title = $1
    WHERE id = $2 AND user_id = $3;

-- name: UpdateCollectionDescription :execrows
UPDATE collections
    SET description = $1
    WHERE id = $2 AND user_id = $3;

-- name: GetArchivedCollectionByID :one
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE id = $1
  AND user_id = $2
  AND archived_at IS NOT NULL;

-- name: GetAllArchivedCollections :many
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE user_id = $1
  AND archived_at IS NOT NULL
ORDER BY archived_at DESC
LIMIT 20;

-- name: GetAllUnarchivedCollections :many
SELECT id, user_id, title, description, created_at, updated_at, archived_at
FROM collections
WHERE user_id = $1
  AND archived_at IS NULL
ORDER BY archived_at DESC
LIMIT 20;

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
LIMIT 20;
