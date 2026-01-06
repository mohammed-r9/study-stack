-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tokens (
    hash TEXT NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    scope TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    used_at TIMESTAMPTZ DEFAULT NULL,
    revoked_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ NOT NULL,

    CHECK (revoked_at IS NULL OR revoked_at >= created_at)
);

CREATE INDEX IF NOT EXISTS idx_tokens_user_id
    ON sessions(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
-- +goose StatementEnd
