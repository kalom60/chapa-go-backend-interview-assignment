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

-- name: UpdateTransfer :one
UPDATE transfer
SET
    account_name = $2,
    account_number = $3,
    currency = $4,
    amount = $5,
    charge = $6,
    transfer_type = $7,
    chapa_reference = $8,
    bank_code = $9,
    bank_name = $10,
    bank_reference = $11,
    status = $12,
    created_at = $13,
    updated_at = $14
WHERE reference = $1
RETURNING *;

-- name: GetTransferByRef :one
SELECT *
FROM transfer
WHERE reference = $1;
