// Code generated by sqlc. DO NOT EDIT.
// source: contas.sql

package db

import (
	"context"
)

const createConta = `-- name: CreateConta :one
INSERT INTO contas (
  proprietario,
  saldo,
  moeda
) VALUES (
  $1, $2, $3
) RETURNING id, proprietario, saldo, moeda, created_at
`

type CreateContaParams struct {
	Proprietario string `json:"proprietario"`
	Saldo        int64  `json:"saldo"`
	Moeda        string `json:"moeda"`
}

func (q *Queries) CreateConta(ctx context.Context, arg CreateContaParams) (Conta, error) {
	row := q.db.QueryRowContext(ctx, createConta, arg.Proprietario, arg.Saldo, arg.Moeda)
	var i Conta
	err := row.Scan(
		&i.ID,
		&i.Proprietario,
		&i.Saldo,
		&i.Moeda,
		&i.CreatedAt,
	)
	return i, err
}

const deleteConta = `-- name: DeleteConta :exec
DELETE FROM contas WHERE id = $1
`

func (q *Queries) DeleteConta(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteConta, id)
	return err
}

const getConta = `-- name: GetConta :one
SELECT id, proprietario, saldo, moeda, created_at FROM contas
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetConta(ctx context.Context, id int64) (Conta, error) {
	row := q.db.QueryRowContext(ctx, getConta, id)
	var i Conta
	err := row.Scan(
		&i.ID,
		&i.Proprietario,
		&i.Saldo,
		&i.Moeda,
		&i.CreatedAt,
	)
	return i, err
}

const listConta = `-- name: ListConta :many
SELECT id, proprietario, saldo, moeda, created_at FROM contas
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListContaParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListConta(ctx context.Context, arg ListContaParams) ([]Conta, error) {
	rows, err := q.db.QueryContext(ctx, listConta, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Conta
	for rows.Next() {
		var i Conta
		if err := rows.Scan(
			&i.ID,
			&i.Proprietario,
			&i.Saldo,
			&i.Moeda,
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

const updateConta = `-- name: UpdateConta :one
UPDATE contas
SET saldo = $2
WHERE id = $1
RETURNING id, proprietario, saldo, moeda, created_at
`

type UpdateContaParams struct {
	ID    int64 `json:"id"`
	Saldo int64 `json:"saldo"`
}

func (q *Queries) UpdateConta(ctx context.Context, arg UpdateContaParams) (Conta, error) {
	row := q.db.QueryRowContext(ctx, updateConta, arg.ID, arg.Saldo)
	var i Conta
	err := row.Scan(
		&i.ID,
		&i.Proprietario,
		&i.Saldo,
		&i.Moeda,
		&i.CreatedAt,
	)
	return i, err
}