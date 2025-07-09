-- name: CreateTransfer :one
INSERT INTO transfer (
    account_name, account_number, currency, amount, charge, transfer_type, chapa_reference,
    bank_code, bank_name, bank_reference, status, reference, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: CountTransfers :one
SELECT COUNT(*) FROM transfer;

-- name: GetAllTransfers :many
SELECT *
FROM transfer
LIMIT $1 OFFSET $2;
