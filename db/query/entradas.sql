-- name: CreateEntradas :one
INSERT INTO entradas (
  id_conta,
  quantidade
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntradas :one
SELECT * FROM entradas
WHERE id = $1 LIMIT 1;

-- name: ListEntradas :many
SELECT * FROM entradas
WHERE id_conta = $1
ORDER BY id
LIMIT $2
OFFSET $3;