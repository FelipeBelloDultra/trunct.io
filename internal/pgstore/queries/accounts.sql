-- name: CreateAccount :one
INSERT INTO accounts (name, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id;

-- name: FindAccountByEmail :one
SELECT id, name, email, password_hash
FROM accounts
WHERE email = $1;

-- name: FindAccountByID :one
SELECT id, name, email, password_hash
FROM accounts
WHERE id = $1;

