-- name: CreateURL :one
INSERT INTO urls (original_url, code, owner_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: FindByCode :one
SELECT id, code, owner_id, original_url, clicks, created_at, updated_at
FROM urls
WHERE code = $1;
