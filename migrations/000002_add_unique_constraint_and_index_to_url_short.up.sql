-- Add a unique constraint to the url_short column
ALTER TABLE urls_short_and_target
ADD CONSTRAINT unique_url_short UNIQUE (url_short);

-- Add an index to the url_short column
CREATE INDEX idx_url_short ON urls_short_and_target(url_short);