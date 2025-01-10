CREATE TABLE IF NOT EXISTS urls_short_and_target (
    id SERIAL PRIMARY KEY,
    url_short TEXT NOT NULL,
    url_target TEXT NOT NULL
);