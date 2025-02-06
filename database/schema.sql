CREATE TABLE IF NOT EXISTS "urls" (
    original_url TEXT PRIMARY KEY NOT NULL,
    shortened_url TEXT NOT NULL,
    clicks INTEGER DEFAULT 0,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uniq_original_url UNIQUE (original_url)
);



