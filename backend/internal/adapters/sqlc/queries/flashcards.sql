-- name: CreateFlashcard :exec
INSERT INTO flashcards (id, material_id, front, back)
SELECT $1, $2, $3, $4
WHERE EXISTS (
    SELECT 1
    FROM materials m
    JOIN collections c ON c.id = m.collection_id
    WHERE m.id = $2 AND c.user_id = $5
);

-- name: getOldestFlashcard :one
SELECT f.*
FROM flashcards f
JOIN materials m ON m.id = f.material_id
JOIN collections c ON c.id = m.collection_id
WHERE c.user_id = $1
ORDER BY f.last_used ASC
LIMIT 1;

-- name: useFlashcard :execrows
UPDATE flashcards
	SET last_used = NOW()
	WHERE id = $1;
