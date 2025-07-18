// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: transfer.sql

package gen

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const CountTransfers = `-- name: CountTransfers :one
SELECT COUNT(*) FROM transfer
`

func (q *Queries) CountTransfers(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, CountTransfers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateTransfer = `-- name: CreateTransfer :one
INSERT INTO transfer (
    account_name, account_number, currency, amount, charge, transfer_type, chapa_reference,
    bank_code, bank_name, bank_reference, status, reference, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING account_name, account_number, currency, amount, charge, transfer_type, chapa_reference, bank_code, bank_name, bank_reference, status, reference, created_at, updated_at
`

type CreateTransferParams struct {
	AccountName    string
	AccountNumber  string
	Currency       string
	Amount         float64
	Charge         float64
	TransferType   string
	ChapaReference string
	BankCode       int32
	BankName       string
	BankReference  pgtype.Text
	Status         string
	Reference      pgtype.Text
	CreatedAt      pgtype.Timestamp
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, CreateTransfer,
		arg.AccountName,
		arg.AccountNumber,
		arg.Currency,
		arg.Amount,
		arg.Charge,
		arg.TransferType,
		arg.ChapaReference,
		arg.BankCode,
		arg.BankName,
		arg.BankReference,
		arg.Status,
		arg.Reference,
		arg.CreatedAt,
	)
	var i Transfer
	err := row.Scan(
		&i.AccountName,
		&i.AccountNumber,
		&i.Currency,
		&i.Amount,
		&i.Charge,
		&i.TransferType,
		&i.ChapaReference,
		&i.BankCode,
		&i.BankName,
		&i.BankReference,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetAllTransfers = `-- name: GetAllTransfers :many
SELECT account_name, account_number, currency, amount, charge, transfer_type, chapa_reference, bank_code, bank_name, bank_reference, status, reference, created_at, updated_at
FROM transfer
LIMIT $1 OFFSET $2
`

type GetAllTransfersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllTransfers(ctx context.Context, arg GetAllTransfersParams) ([]Transfer, error) {
	rows, err := q.db.Query(ctx, GetAllTransfers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.AccountName,
			&i.AccountNumber,
			&i.Currency,
			&i.Amount,
			&i.Charge,
			&i.TransferType,
			&i.ChapaReference,
			&i.BankCode,
			&i.BankName,
			&i.BankReference,
			&i.Status,
			&i.Reference,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetTransferByRef = `-- name: GetTransferByRef :one
SELECT account_name, account_number, currency, amount, charge, transfer_type, chapa_reference, bank_code, bank_name, bank_reference, status, reference, created_at, updated_at
FROM transfer
WHERE reference = $1
`

func (q *Queries) GetTransferByRef(ctx context.Context, reference pgtype.Text) (Transfer, error) {
	row := q.db.QueryRow(ctx, GetTransferByRef, reference)
	var i Transfer
	err := row.Scan(
		&i.AccountName,
		&i.AccountNumber,
		&i.Currency,
		&i.Amount,
		&i.Charge,
		&i.TransferType,
		&i.ChapaReference,
		&i.BankCode,
		&i.BankName,
		&i.BankReference,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const UpdateTransfer = `-- name: UpdateTransfer :one
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
RETURNING account_name, account_number, currency, amount, charge, transfer_type, chapa_reference, bank_code, bank_name, bank_reference, status, reference, created_at, updated_at
`

type UpdateTransferParams struct {
	Reference      pgtype.Text
	AccountName    string
	AccountNumber  string
	Currency       string
	Amount         float64
	Charge         float64
	TransferType   string
	ChapaReference string
	BankCode       int32
	BankName       string
	BankReference  pgtype.Text
	Status         string
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, UpdateTransfer,
		arg.Reference,
		arg.AccountName,
		arg.AccountNumber,
		arg.Currency,
		arg.Amount,
		arg.Charge,
		arg.TransferType,
		arg.ChapaReference,
		arg.BankCode,
		arg.BankName,
		arg.BankReference,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Transfer
	err := row.Scan(
		&i.AccountName,
		&i.AccountNumber,
		&i.Currency,
		&i.Amount,
		&i.Charge,
		&i.TransferType,
		&i.ChapaReference,
		&i.BankCode,
		&i.BankName,
		&i.BankReference,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
