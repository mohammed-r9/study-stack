-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS image_labels (
    id UUID PRIMARY KEY,
    image_id UUID NOT NULL REFERENCES image_flashcards(id) ON DELETE CASCADE,
	x_start FLOAT NOT NULL,
	x_end FLOAT NOT NULL,
	y_start FLOAT NOT NULL,
	y_end FLOAT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_image_labels_image_id 
    ON image_labels(image_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE image_labels;
-- +goose StatementEnd
