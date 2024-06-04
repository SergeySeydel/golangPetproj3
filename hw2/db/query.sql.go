// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const getCard = `-- name: GetCard :one
SELECT id, time, date, title, client, "user" FROM cards
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCard(ctx context.Context, id int64) (Card, error) {
	row := q.db.QueryRow(ctx, getCard, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Time,
		&i.Date,
		&i.Title,
		&i.Client,
		&i.User,
	)
	return i, err
}
