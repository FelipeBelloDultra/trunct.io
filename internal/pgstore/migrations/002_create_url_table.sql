-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS urls (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  owner_id UUID NOT NULL REFERENCES accounts (id),
  original_url VARCHAR(255) NOT NULL,
  short_url VARCHAR(255) UNIQUE NOT NULL,
  clicks INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABLE IF EXISTS urls;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
