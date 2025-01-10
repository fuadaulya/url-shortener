-- Remove the unique constraint from the url_short column
ALTER TABLE urls_short_and_target
DROP CONSTRAINT IF EXISTS unique_url_short;

-- Remove the index from the url_short column
DROP INDEX IF EXISTS idx_url_short;