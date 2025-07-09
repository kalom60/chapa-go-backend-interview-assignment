package transaction

import (
	"context"

	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, tx Transaction) error
	UpdateTransaction(ctx context.Context, tx Transaction) error
	GetTransactionByRef(ctx context.Context, ref string) (Transaction, error)
	GetAllTransactions(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransactions[Transaction], error)
}
