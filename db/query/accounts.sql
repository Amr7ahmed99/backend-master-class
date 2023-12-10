-- name: GetAccount :one
SELECT * FROM "accounts"
WHERE id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM "accounts"
ORDER BY id
LIMIT $1;

-- name: CreateAccount :one
INSERT INTO "accounts" (
  owner, balance, currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM "accounts"
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts
SET balance= $2, currency= $3
WHERE id= $1
RETURNING *;