-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS lectures (
    id UUID PRIMARY KEY,
    material_id UUID NOT NULL REFERENCES materials(id) ON DELETE CASCADE,
    title TEXT NOT NULL,

    file_key UUID NOT NULL UNIQUE,
    file_size BIGINT NOT NULL CHECK (file_size > 0),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    archived_at TIMESTAMPTZ

);

CREATE INDEX IF NOT EXISTS idx_lectures_material_id 
    ON lectures(material_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lectures;
-- +goose StatementEnd
