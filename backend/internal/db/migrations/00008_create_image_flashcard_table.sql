-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS image_flashcards (
    id UUID PRIMARY KEY,
    material_id UUID NOT NULL REFERENCES materials(id) ON DELETE CASCADE,
	title TEXT NOT NULL,
	image_key TEXT NOT NULL UNIQUE,
    image_size BIGINT NOT NULL CHECK (image_size > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_used TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_image_flashcards_material_id 
    ON image_flashcards(material_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE image_flashcards;
-- +goose StatementEnd
