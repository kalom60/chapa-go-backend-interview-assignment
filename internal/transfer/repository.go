package transfer

import (
	"context"

	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

type TransferRepo interface {
	CreateTransfer(ctx context.Context, transfer Transfer) error
	GetAllTransfers(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransfers[Transfer], error)
}
