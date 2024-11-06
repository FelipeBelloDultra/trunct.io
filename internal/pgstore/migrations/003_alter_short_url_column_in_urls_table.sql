-- Write your migrate up statements here
ALTER TABLE urls
RENAME COLUMN short_url to code;
---- create above / drop below ----
ALTER TABLE urls
RENAME COLUMN code to short_url;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
