package bank

import (
	"context"

	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

type BankRepo interface {
	GetAllBanks(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseBanks[Bank], error)
	GetBankByBankID(ctx context.Context, bankID int) (Bank, error)
}
