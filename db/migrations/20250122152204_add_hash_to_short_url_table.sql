-- +goose Up
-- +goose StatementBegin
ALTER TABLE short_urls ADD COLUMN hash TEXT NOT NULL UNIQUE DEFAULT 'default_hash';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE short_urls DROP COLUMN hash;
-- +goose StatementEnd
