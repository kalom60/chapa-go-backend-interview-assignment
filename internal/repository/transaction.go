package repository

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/repository/gen"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/transaction"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

func (repo *Store) CreateTransaction(ctx context.Context, tx transaction.Transaction) error {
	_, err := repo.queries.CreateTransaction(ctx, gen.CreateTransactionParams{
		Status: tx.Status,
		RefID:  tx.RefID,
		Type:   tx.Type,
		CreatedAt: pgtype.Timestamp{
			Time:  tx.CreatedAt,
			Valid: true,
		},
		Currency: tx.Currency,
		Amount:   tx.Amount,
		Charge:   tx.Charge,
		TransID: pgtype.Text{
			String: tx.TransID,
			Valid:  true,
		},
		PaymentMethod: tx.PaymentMethod,
		CustomerID:    tx.CustomerID,
		CustomerFirstName: pgtype.Text{
			String: tx.CustomerFirstname,
			Valid:  true,
		},
		CustomerLastName: pgtype.Text{
			String: tx.CustomerLastname,
			Valid:  true,
		},
		CustomerMobile: pgtype.Text{
			String: tx.CustomerMobile,
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *Store) UpdateTransaction(ctx context.Context, tx transaction.Transaction) error {
	_, err := repo.queries.UpdateTransaction(ctx, gen.UpdateTransactionParams{
		Status: tx.Status,
		RefID:  tx.RefID,
		Type:   tx.Type,
		CreatedAt: pgtype.Timestamp{
			Time:  tx.CreatedAt,
			Valid: true,
		},
		Currency: tx.Currency,
		Amount:   tx.Amount,
		Charge:   tx.Charge,
		TransID: pgtype.Text{
			String: tx.TransID,
			Valid:  true,
		},
		PaymentMethod: tx.PaymentMethod,
		CustomerID:    tx.CustomerID,
		CustomerFirstName: pgtype.Text{
			String: tx.CustomerFirstname,
			Valid:  true,
		},
		CustomerLastName: pgtype.Text{
			String: tx.CustomerLastname,
			Valid:  true,
		},
		CustomerMobile: pgtype.Text{
			String: tx.CustomerMobile,
			Valid:  true,
		},
	})

	return err
}

func (repo *Store) GetTransactionByRef(ctx context.Context, ref string) (transaction.Transaction, error) {
	row, err := repo.queries.GetTransactionByRef(ctx, ref)
	if err != nil {
		return transaction.Transaction{}, nil
	}

	tr := transaction.Transaction{
		Status:            row.Status,
		RefID:             row.RefID,
		Type:              row.Type,
		CreatedAt:         row.CreatedAt.Time,
		Currency:          row.Currency,
		Amount:            row.Amount,
		Charge:            row.Charge,
		TransID:           row.TransID.String,
		PaymentMethod:     row.PaymentMethod,
		CustomerID:        row.CustomerID,
		CustomerFirstname: row.CustomerFirstName.String,
		CustomerLastname:  row.CustomerLastName.String,
		CustomerMobile:    row.CustomerMobile.String,
	}

	return tr, nil
}

func (repo *Store) GetAllTransactions(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransactions[transaction.Transaction], error) {
	count, err := repo.queries.CountTransactions(ctx)
	if err != nil {
		return utils.PaginatedResponseTransactions[transaction.Transaction]{}, err
	}

	rows, err := repo.queries.GetAllTransactions(ctx, gen.GetAllTransactionsParams{
		Limit:  int32(filter.PageSize),
		Offset: int32((filter.Page - 1) * filter.PageSize),
	})
	if err != nil {
		return utils.PaginatedResponseTransactions[transaction.Transaction]{}, err
	}

	txs := make([]transaction.Transaction, 0, len(rows))
	for _, row := range rows {
		txs = append(txs, transaction.Transaction{
			Status:            row.Status,
			RefID:             row.RefID,
			Type:              row.Type,
			CreatedAt:         row.CreatedAt.Time,
			Currency:          row.Currency,
			Amount:            row.Amount,
			Charge:            row.Charge,
			TransID:           row.TransID.String,
			PaymentMethod:     row.PaymentMethod,
			CustomerID:        row.CustomerID,
			CustomerFirstname: row.CustomerFirstName.String,
			CustomerLastname:  row.CustomerLastName.String,
			CustomerMobile:    row.CustomerMobile.String,
		})
	}

	response := utils.PaginatedResponseTransactions[transaction.Transaction]{
		Transactions: txs,
		Meta: utils.Meta{
			ItemsPerPage: filter.PageSize,
			TotalItems:   count,
			CurrentPage:  filter.Page,
			TotalPages:   int(math.Ceil(float64(count) / float64(filter.PageSize))),
		},
	}

	return response, nil
}
