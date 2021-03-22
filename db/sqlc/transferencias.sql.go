// Code generated by sqlc. DO NOT EDIT.
// source: transferencias.sql

package db

import (
	"context"
)

const createTransferencias = `-- name: CreateTransferencias :one
INSERT INTO transferencias (
  emissor_id_conta,
  receptor_id_conta,
  quantidade
) VALUES (
  $1, $2, $3
) RETURNING id, emissor_id_conta, receptor_id_conta, quantidade, created_at
`

type CreateTransferenciasParams struct {
	EmissorIDConta  int64  `json:"emissor_id_conta"`
	ReceptorIDConta int64  `json:"receptor_id_conta"`
	Quantidade      string `json:"quantidade"`
}

func (q *Queries) CreateTransferencias(ctx context.Context, arg CreateTransferenciasParams) (Transferencia, error) {
	row := q.db.QueryRowContext(ctx, createTransferencias, arg.EmissorIDConta, arg.ReceptorIDConta, arg.Quantidade)
	var i Transferencia
	err := row.Scan(
		&i.ID,
		&i.EmissorIDConta,
		&i.ReceptorIDConta,
		&i.Quantidade,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferencias = `-- name: GetTransferencias :one
SELECT id, emissor_id_conta, receptor_id_conta, quantidade, created_at FROM transferencias
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransferencias(ctx context.Context, id int64) (Transferencia, error) {
	row := q.db.QueryRowContext(ctx, getTransferencias, id)
	var i Transferencia
	err := row.Scan(
		&i.ID,
		&i.EmissorIDConta,
		&i.ReceptorIDConta,
		&i.Quantidade,
		&i.CreatedAt,
	)
	return i, err
}

const listTransferencias = `-- name: ListTransferencias :many
SELECT id, emissor_id_conta, receptor_id_conta, quantidade, created_at FROM transferencias
WHERE 
    emissor_id_conta = $1 OR
    receptor_id_conta = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransferenciasParams struct {
	EmissorIDConta  int64 `json:"emissor_id_conta"`
	ReceptorIDConta int64 `json:"receptor_id_conta"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *Queries) ListTransferencias(ctx context.Context, arg ListTransferenciasParams) ([]Transferencia, error) {
	rows, err := q.db.QueryContext(ctx, listTransferencias,
		arg.EmissorIDConta,
		arg.ReceptorIDConta,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transferencia
	for rows.Next() {
		var i Transferencia
		if err := rows.Scan(
			&i.ID,
			&i.EmissorIDConta,
			&i.ReceptorIDConta,
			&i.Quantidade,
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