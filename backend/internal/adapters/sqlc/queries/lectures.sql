-- name: CreateLecture :exec
INSERT INTO lectures (id, material_id, title, file_key, file_size)
SELECT $1, $2, $3, $4, $5
WHERE EXISTS (
    SELECT 1
    FROM materials m
    JOIN collections c ON c.id = m.collection_id
    WHERE m.id = $2 AND c.user_id = $6
);

-- name: GetLectureByID :one
SELECT l.*
FROM lectures l
JOIN materials m ON m.id = l.material_id
JOIN collections c ON c.id = m.collection_id
WHERE l.id = $1 AND c.user_id = $2;

-- name: GetLectureFileKey :one
SELECT l.file_key, l.id
FROM lectures l
JOIN materials m ON m.id = l.material_id
JOIN collections c ON c.id = m.collection_id
WHERE l.id = $1 AND c.user_id = $2;

-- name: UpdateLectureTitle :execrows
UPDATE lectures l
SET title = $3, updated_at = CURRENT_TIMESTAMP
FROM materials m
JOIN collections c ON c.id = m.collection_id
WHERE l.id = $1 AND l.material_id = m.id AND c.user_id = $2;

-- name: DeleteLecture :execrows
DELETE FROM lectures l
USING materials m, collections c
WHERE l.id = $1 AND l.material_id = m.id
  AND m.collection_id = c.id
  AND c.user_id = $2;

-- name: ArchiveLecture :execrows
UPDATE lectures l
SET archived_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
FROM materials m
JOIN collections c ON c.id = m.collection_id
WHERE l.id = $1 AND l.material_id = m.id AND c.user_id = $2;

-- name: UnarchiveLecture :execrows
UPDATE lectures l
SET archived_at = NULL, updated_at = CURRENT_TIMESTAMP
FROM materials m
JOIN collections c ON c.id = m.collection_id
WHERE l.id = $1 AND l.material_id = m.id AND c.user_id = $2;

-- name: ListLectures :many
SELECT l.id,
       l.material_id,
       l.title,
       l.file_size,
       l.created_at,
       l.updated_at,
       l.archived_at
FROM lectures l
JOIN materials m ON m.id = l.material_id
JOIN collections c ON c.id = m.collection_id
WHERE c.user_id = @user_id
  AND m.id = @material_id
  AND l.id < @last_seen_lecture_id
ORDER BY l.id DESC
LIMIT 20;


-- name: ListActiveLectures :many
SELECT l.*
FROM lectures l
JOIN materials m ON m.id = l.material_id
JOIN collections c ON c.id = m.collection_id
WHERE c.user_id = $1 AND l.archived_at IS NULL
ORDER BY l.created_at DESC, l.id DESC
LIMIT 20 OFFSET $2;


-- name: ListArchivedLectures :many
SELECT l.*
FROM lectures l
JOIN materials m ON m.id = l.material_id
JOIN collections c ON c.id = m.collection_id
WHERE c.user_id = $1 AND l.archived_at IS NOT NULL
ORDER BY l.created_at DESC, l.id DESC
LIMIT 20 OFFSET $2;
