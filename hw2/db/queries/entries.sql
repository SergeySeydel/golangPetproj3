entries.sql
-- name: CreateEntry :one
INSERT INTO entries (content) VALUES ($1) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1;

-- name: ListEntries :many
SELECT * FROM entries ORDER BY created_at DESC;

-- name: UpdateEntry :one
UPDATE entries SET content = $1 WHERE id = $2 RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;
