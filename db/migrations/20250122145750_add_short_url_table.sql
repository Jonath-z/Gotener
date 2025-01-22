-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    short_urls
    (
        id SERIAL PRIMARY KEY,
        long_url TEXT NOT NULL,
        short_url TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS short_urls;
-- +goose StatementEnd
