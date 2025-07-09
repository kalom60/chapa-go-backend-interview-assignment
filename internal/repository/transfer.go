package repository

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/repository/gen"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/transfer"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

func (repo *Store) CreateTransfer(ctx context.Context, transfer transfer.Transfer) error {
	_, err := repo.queries.CreateTransfer(ctx, gen.CreateTransferParams{
		AccountName:    transfer.AccountName,
		AccountNumber:  transfer.AccountNumber,
		Currency:       transfer.Currency,
		Amount:         float64(transfer.Amount),
		Charge:         float64(transfer.Charge),
		TransferType:   transfer.TransferType,
		ChapaReference: transfer.ChapaReference,
		BankCode:       int32(transfer.BankCode),
		BankName:       transfer.BankName,
		BankReference: pgtype.Text{
			String: transfer.BankReference,
			Valid:  true,
		},
		Status: transfer.Status,
		Reference: pgtype.Text{
			String: transfer.Reference,
			Valid:  true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  transfer.CreatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *Store) GetAllTransfers(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransfers[transfer.Transfer], error) {
	count, err := repo.queries.CountBanks(ctx)
	if err != nil {
		return utils.PaginatedResponseTransfers[transfer.Transfer]{}, err
	}

	rows, err := repo.queries.GetAllTransfers(ctx, gen.GetAllTransfersParams{
		Limit:  int32(filter.PageSize),
		Offset: int32((filter.Page - 1) * filter.PageSize),
	})
	if err != nil {
		return utils.PaginatedResponseTransfers[transfer.Transfer]{}, err
	}

	banks := make([]transfer.Transfer, 0, len(rows))
	for _, row := range rows {
		banks = append(banks, transfer.Transfer{
			AccountName:    row.AccountName,
			AccountNumber:  row.AccountNumber,
			Currency:       row.Currency,
			Amount:         float64(row.Amount),
			Charge:         float64(row.Charge),
			TransferType:   row.TransferType,
			ChapaReference: row.ChapaReference,
			BankCode:       int(row.BankCode),
			BankName:       row.BankName,
			BankReference:  row.BankReference.String,
			Status:         row.Status,
			Reference:      row.Reference.String,
			CreatedAt:      row.CreatedAt.Time,
			UpdatedAt:      row.UpdatedAt.Time,
		})
	}

	response := utils.PaginatedResponseTransfers[transfer.Transfer]{
		Transfers: banks,
		Meta: utils.Meta{
			ItemsPerPage: filter.PageSize,
			TotalItems:   count,
			CurrentPage:  filter.Page,
			TotalPages:   int(math.Ceil(float64(count) / float64(filter.PageSize))),
		},
	}

	return response, nil
}
func (repo *Store) UpdateTransfer(ctx context.Context, tr transfer.Transfer) error {
	_, err := repo.queries.UpdateTransfer(ctx, gen.UpdateTransferParams{
		Reference: pgtype.Text{
			String: tr.Reference,
			Valid:  true,
		},
		AccountName:    tr.AccountName,
		AccountNumber:  tr.AccountNumber,
		Currency:       tr.Currency,
		Amount:         tr.Amount,
		Charge:         tr.Charge,
		TransferType:   tr.TransferType,
		ChapaReference: tr.ChapaReference,
		BankCode:       int32(tr.BankCode),
		BankName:       tr.BankName,
		BankReference: pgtype.Text{
			String: tr.BankReference,
			Valid:  true,
		},
		Status: tr.Status,
		CreatedAt: pgtype.Timestamp{
			Time:  tr.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  tr.UpdatedAt,
			Valid: true,
		},
	})

	return err
}

func (repo *Store) GetTransferByRef(ctx context.Context, ref string) (transfer.Transfer, error) {
	row, err := repo.queries.GetTransferByRef(ctx, pgtype.Text{
		String: ref,
		Valid:  true,
	})
	if err != nil {
		return transfer.Transfer{}, nil
	}

	tr := transfer.Transfer{
		AccountName:    row.AccountName,
		AccountNumber:  row.AccountNumber,
		Currency:       row.Currency,
		Amount:         float64(row.Amount),
		Charge:         float64(row.Charge),
		TransferType:   row.TransferType,
		ChapaReference: row.ChapaReference,
		BankCode:       int(row.BankCode),
		BankName:       row.BankName,
		BankReference:  row.BankReference.String,
		Status:         row.Status,
		Reference:      row.Reference.String,
		CreatedAt:      row.CreatedAt.Time,
		UpdatedAt:      row.UpdatedAt.Time,
	}

	return tr, nil
}
