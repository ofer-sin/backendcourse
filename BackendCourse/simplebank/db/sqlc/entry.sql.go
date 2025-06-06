// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entries, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many


SELECT id, account_id, amount, created_at FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

// This query retrieves a list of entries from the "entries" table that belong to a specific account (filtered by account_id).
// The results are ordered by the "id" column in ascending order.
// The "LIMIT $2" clause restricts the number of rows returned to the value specified by the second parameter.
// The "OFFSET $3" clause skips the first $3 rows, allowing for pagination of results.
// Key Points:
// account_id = $1: Filters rows based on the provided account ID.
// ORDER BY id: Ensures the results are sorted by the id column in ascending order.
// LIMIT $2: Limits the number of rows returned to the value of $2.
// OFFSET $3: Skips the first $3 rows, useful for implementing pagination.
// This query is commonly used in applications to fetch a subset of data for a specific account, often for displaying paginated results in a UI.
func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entries{}
	for rows.Next() {
		var i Entries
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
