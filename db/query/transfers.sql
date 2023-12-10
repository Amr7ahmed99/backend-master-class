-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1;

-- name: CreateTransfer :one
INSERT INTO transfers (
  owner, from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfer :one
UPDATE transfers 
SET owner= $2, from_account_id= $3, 
to_account_id= $4, amount= $5
WHERE id= $1
RETURNING *;