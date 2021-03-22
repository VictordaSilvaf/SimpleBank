-- name: CreateConta :one
INSERT INTO contas (
  proprietario,
  saldo,
  moeda
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetConta :one
SELECT * FROM contas
WHERE id = $1 LIMIT 1;

-- name: ListConta :many
SELECT * FROM contas
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateConta :one
UPDATE contas
SET saldo = $2
WHERE id = $1
RETURNING *;

-- name: DeleteConta :exec
DELETE FROM contas WHERE id = $1;