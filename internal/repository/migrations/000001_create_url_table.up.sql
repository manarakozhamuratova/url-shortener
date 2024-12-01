CREATE TABLE IF NOT EXISTS short_url (
    id SERIAL PRIMARY KEY,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    short_url VARCHAR(255) NOT NULL UNIQUE,
    long_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    expires_at TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX idx_unique_short_url ON short_url (short_url);