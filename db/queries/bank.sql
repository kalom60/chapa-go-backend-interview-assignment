-- name: CountBanks :one
SELECT COUNT(*) FROM bank;

-- name: GetAllBanks :many
SELECT *
FROM bank
LIMIT $1 OFFSET $2;

-- name: GetBankByBankID :one
SELECT *
FROM bank
WHERE bank_id = $1;
