-- name: CreateTransaction :one
INSERT INTO transaction(
    status, ref_id, type, created_at, currency, amount, charge,
    trans_id, payment_method, customer_id, customer_first_name, customer_last_name,customer_mobile
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: CountTransactions :one
SELECT COUNT(*) FROM transaction;

-- name: GetAllTransactions :many
SELECT *
FROM transaction
LIMIT $1 OFFSET $2;

-- name: UpdateTransaction :one
UPDATE transaction
SET
    status= $2,
    ref_id= $3,
    type= $4,
    created_at= $5,
    currency= $6,
    amount= $7,
    charge= $8,
    trans_id= $9,
    payment_method= $10,
    customer_id= $11,
    customer_first_name= $12,
    customer_last_name= $13,
    customer_mobile= $14
WHERE ref_id= $1
RETURNING *;

-- name: GetTransactionByRef :one
SELECT *
FROM transaction
WHERE ref_id= $1;
