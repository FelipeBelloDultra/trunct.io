-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS accounts (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password_hash BYTEA NOT NULL
);
---- create above / drop below ----
DROP TABLE IF EXISTS accounts;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
