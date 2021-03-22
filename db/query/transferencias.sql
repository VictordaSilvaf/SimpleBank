-- name: CreateTransferencias :one
INSERT INTO transferencias (
  emissor_id_conta,
  receptor_id_conta,
  quantidade
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransferencias :one
SELECT * FROM transferencias
WHERE id = $1 LIMIT 1;

-- name: ListTransferencias :many
SELECT * FROM transferencias
WHERE 
    emissor_id_conta = $1 OR
    receptor_id_conta = $2
ORDER BY id
LIMIT $3
OFFSET $4;