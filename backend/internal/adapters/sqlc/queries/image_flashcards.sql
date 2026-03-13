-- name: createImageFlashcard :one
INSERT INTO image_flashcards (id, material_id, title, image_key, image_size)
SELECT $1, $2, $3, $4, $5
WHERE EXISTS (
    SELECT 1
    FROM materials m
    JOIN collections c ON c.id = m.collection_id
    WHERE m.id = $2 AND c.user_id = $6
)
RETURNING image_flashcards.*;
