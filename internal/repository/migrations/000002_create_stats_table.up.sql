CREATE TABLE IF NOT EXISTS stats (
    url_id        INT           NOT NULL,
    visited_count INT           NOT NULL DEFAULT 0,
    visited_at    TIMESTAMPTZ   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_stats PRIMARY KEY (url_id),    -- Ensure each url_id is unique
    CONSTRAINT fk_url FOREIGN KEY (url_id)        -- Foreign key constraint
        REFERENCES short_url(id)                   -- Assuming `id` is the primary key of `short_url`
        ON DELETE CASCADE                          -- Optional: Delete stats if the associated short_url is deleted
);
