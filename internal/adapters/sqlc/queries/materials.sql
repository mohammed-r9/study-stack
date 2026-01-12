-- name: InsertMaterial :exec
INSERT INTO materials (id, collection_id, title)
SELECT $1, $2, $3
WHERE (
    SELECT COUNT(*) 
    FROM materials 
    WHERE collection_id = $2
) < 20
AND EXISTS (
    SELECT 1 
    FROM collections 
    WHERE id = $2 AND user_id = $4
);

-- name: GetMaterialByID :one
SELECT 
    m.id, 
    m.collection_id, 
    m.title, 
    m.created_at, 
    m.updated_at, 
    m.archived_at
FROM materials m
JOIN collections c ON m.collection_id = c.id
WHERE m.id = $1 
  AND c.user_id = $2
  AND m.archived_at IS NULL;

-- name: GetAllUnarchivedMaterialsInCollection :many
SELECT 
    m.id, 
    m.collection_id, 
    m.title, 
    m.created_at, 
    m.updated_at, 
    m.archived_at
FROM materials m
JOIN collections c ON m.collection_id = c.id
WHERE c.user_id = $1 
  AND m.collection_id = $2
  AND m.archived_at IS NULL
LIMIT 20;

-- name: GetAllArchivedMaterialsInCollection :many
SELECT 
    m.id, 
    m.collection_id, 
    m.title, 
    m.created_at, 
    m.updated_at, 
    m.archived_at
FROM materials m
JOIN collections c ON m.collection_id = c.id
WHERE c.user_id = $1 
  AND m.collection_id = $2
  AND m.archived_at IS NOT NULL
LIMIT 20;

-- name: GetAllMaterialsInCollection :many
SELECT 
    m.id, 
    m.collection_id, 
    m.title, 
    m.created_at, 
    m.updated_at, 
    m.archived_at
FROM materials m
JOIN collections c ON m.collection_id = c.id
WHERE c.user_id = $1 
  AND m.collection_id = $2
LIMIT 20;

-- name: ArchiveMaterial :execrows
UPDATE materials
SET archived_at = CURRENT_TIMESTAMP
WHERE materials.id = $1 
  AND materials.archived_at IS NULL
  AND EXISTS (
      SELECT 1 
      FROM collections 
      WHERE collections.id = materials.collection_id 
        AND collections.user_id = $2
  );

-- name: UnarchiveMaterial :execrows
UPDATE materials
SET archived_at = NULL
WHERE materials.id = $1 
  AND materials.archived_at IS NOT NULL
  AND EXISTS (
      SELECT 1 
      FROM collections 
      WHERE collections.id = materials.collection_id 
        AND collections.user_id = $2
  )
  AND (
      SELECT COUNT(*) 
      FROM materials 
      WHERE collection_id = materials.collection_id 
        AND archived_at IS NULL
  ) < 20;
