-- name: CreateTransfer :one
INSERT INTO transfers (from_entry_id, to_entry_id, amount) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = $1;

-- name: ListTransfers :many
SELECT * FROM transfers ORDER BY created_at DESC;

-- name: UpdateTransfer :one
UPDATE transfers SET from_entry_id = $1, to_entry_id = $2, amount = $3 WHERE id = $4 RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;
