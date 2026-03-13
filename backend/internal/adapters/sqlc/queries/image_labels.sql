-- name: insertImageLabel :one
INSERT INTO image_labels (id, image_id, x_start, x_end, y_start, y_end) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING image_labels.*;
