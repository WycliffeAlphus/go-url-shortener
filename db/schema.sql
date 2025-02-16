CREATE TABLE IF NOT EXISTS "urls" (
    original_url TEXT PRIMARY KEY NOT NULL,
    shortened_url TEXT NOT NULL,
    CONSTRAINT uniq_original_url UNIQUE (original_url)
);




