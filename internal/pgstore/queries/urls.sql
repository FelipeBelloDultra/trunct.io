-- name: CreateURL :one
INSERT INTO urls (original_url, code, owner_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: FindURLByCode :one
SELECT id, code, owner_id, original_url, clicks, created_at, updated_at
FROM urls
WHERE code = $1;

-- name: UpdateURL :one
UPDATE urls
SET original_url = $2, clicks = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, code, owner_id, original_url, clicks, created_at, updated_at;
