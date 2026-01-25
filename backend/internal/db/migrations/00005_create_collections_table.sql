-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS materials (
    id UUID NOT NULL PRIMARY KEY,
    collection_id UUID NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    archived_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_materials_collection_id 
    ON materials(collection_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE materials;
-- +goose StatementEnd
