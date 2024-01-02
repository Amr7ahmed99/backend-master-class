// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: currencies.sql

package db

import (
	"context"
)

const getCurrency = `-- name: GetCurrency :one
SELECT id, currency FROM "currencies"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCurrency(ctx context.Context, id int64) (Currency, error) {
	row := q.queryRow(ctx, q.getCurrencyStmt, getCurrency, id)
	var i Currency
	err := row.Scan(&i.ID, &i.Currency)
	return i, err
}

const listCurrencies = `-- name: ListCurrencies :many
SELECT id, currency FROM "currencies"
ORDER BY id ASC
`

func (q *Queries) ListCurrencies(ctx context.Context) ([]Currency, error) {
	rows, err := q.query(ctx, q.listCurrenciesStmt, listCurrencies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Currency{}
	for rows.Next() {
		var i Currency
		if err := rows.Scan(&i.ID, &i.Currency); err != nil {
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