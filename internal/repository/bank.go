package repository

import (
	"context"
	"math"

	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/bank"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/repository/gen"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

func (repo *Store) GetAllBanks(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseBanks[bank.Bank], error) {
	count, err := repo.queries.CountBanks(ctx)
	if err != nil {
		return utils.PaginatedResponseBanks[bank.Bank]{}, err
	}

	rows, err := repo.queries.GetAllBanks(ctx, gen.GetAllBanksParams{
		Limit:  int32(filter.PageSize),
		Offset: int32((filter.Page - 1) * filter.PageSize),
	})
	if err != nil {
		return utils.PaginatedResponseBanks[bank.Bank]{}, err
	}

	banks := make([]bank.Bank, 0, len(rows))
	for _, row := range rows {
		banks = append(banks, bank.Bank{
			BankID:        int(row.BankID),
			Slug:          row.Slug,
			Swift:         row.Swift,
			Name:          row.Name,
			AcctLength:    int(row.AcctLength),
			CountryID:     int(row.CountryID),
			IsMobilemoney: int(row.IsMobilemoney),
			IsActive:      int(row.IsActive),
			Is_Rtgs:       int(row.IsRtgs),
			Active:        int(row.Active),
			Is24hrs:       int(row.Is24hrs),
			CreatedAt:     row.CreatedAt.Time,
			UpdatedAt:     row.UpdatedAt.Time,
			Currency:      row.Currency,
		})
	}

	response := utils.PaginatedResponseBanks[bank.Bank]{
		Banks: banks,
		Meta: utils.Meta{
			ItemsPerPage: filter.PageSize,
			TotalItems:   count,
			CurrentPage:  filter.Page,
			TotalPages:   int(math.Ceil(float64(count) / float64(filter.PageSize))),
		},
	}

	return response, nil
}

func (repo *Store) GetBankByBankID(ctx context.Context, bankID int) (bank.Bank, error) {
	row, err := repo.queries.GetBankByBankID(ctx, int32(bankID))
	if err != nil {
		return bank.Bank{}, err
	}

	bank := bank.Bank{
		BankID:        int(row.BankID),
		Slug:          row.Slug,
		Swift:         row.Swift,
		Name:          row.Name,
		AcctLength:    int(row.AcctLength),
		CountryID:     int(row.CountryID),
		IsMobilemoney: int(row.IsMobilemoney),
		IsActive:      int(row.IsActive),
		Is_Rtgs:       int(row.IsRtgs),
		Active:        int(row.Active),
		Is24hrs:       int(row.Is24hrs),
		CreatedAt:     row.CreatedAt.Time,
		UpdatedAt:     row.UpdatedAt.Time,
		Currency:      row.Currency,
	}

	return bank, nil
}
