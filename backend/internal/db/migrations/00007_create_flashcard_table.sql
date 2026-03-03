-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flashcards (
    id UUID PRIMARY KEY,
    material_id UUID NOT NULL REFERENCES materials(id) ON DELETE CASCADE,
    front TEXT NOT NULL,
    back TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_used TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_flashcards_material_id 
    ON flashcards(material_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE flashcards;
-- +goose StatementEnd
