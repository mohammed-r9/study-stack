-- name: insertImageLabel :one
INSERT INTO image_labels (id, image_id, x_start_percentage, y_start_percentage, width_percentage, height_percentage) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING image_labels.*;
