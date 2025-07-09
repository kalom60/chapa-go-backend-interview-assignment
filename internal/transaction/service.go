package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Chapa-Et/chapa-go"
	"github.com/google/uuid"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/cache"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
	"github.com/redis/go-redis/v9"
)

var (
	ErrDuplicateTransaction = errors.New("duplicate transaction")
	ErrFailedInitiate       error
)

type Service struct {
	txRepo   TransactionRepo
	chapaAPI chapa.API
	redis    cache.RedisCache
}

func NewService(txRepo TransactionRepo, chapaAPI chapa.API, redis cache.RedisCache) *Service {
	return &Service{
		txRepo:   txRepo,
		chapaAPI: chapaAPI,
		redis:    redis,
	}
}

func (s *Service) InitiateTransaction(ctx context.Context, tx *chapa.PaymentRequest) (*chapa.PaymentResponse, error) {
	fmt.Println("TX", tx)
	ref := uuid.New()

	val, err := s.redis.Get(ref.String())
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
	}

	if val != nil {
		return nil, ErrDuplicateTransaction
	}

	ltx, err := createTransaction(tx, ref.String())
	if err != nil {
		return nil, err
	}

	err = s.redis.Set(ref.String(), ltx, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	tx.TransactionRef = ref.String()
	resp, err := s.chapaAPI.PaymentRequest(tx)
	if err != nil {
		return nil, err
	}

	fmt.Println("HERE", resp)

	if resp.Status == "failed" {
		err := s.redis.Delete(ref.String())
		if err != nil {
			return nil, err
		}

		ErrFailedInitiate = errors.New(resp.Message)
		return nil, ErrFailedInitiate
	}

	err = s.txRepo.CreateTransaction(ctx, ltx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) GetTransactions(ctx context.Context) (*chapa.TransactionsResponse, error) {
	return s.chapaAPI.GetTransactions()
}

func (s *Service) GetAllTransactions(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransactions[Transaction], error) {
	return s.txRepo.GetAllTransactions(ctx, filter)
}

func (s *Service) VerifyTransaction(ref string) (*chapa.VerifyResponse, error) {
	return s.chapaAPI.Verify(ref)
}

func (s *Service) HandleWebhook(ctx context.Context, tx Transaction) error {
	if _, err := s.redis.Get(tx.RefID); err != nil {
		if err == redis.Nil {
			_, err := s.txRepo.GetTransactionByRef(ctx, tx.RefID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return err
				}
				return err
			}
		}
		return err
	}

	err := s.txRepo.UpdateTransaction(ctx, tx)
	if err != nil {
		return err
	}

	err = s.redis.Delete(tx.RefID)
	if err != nil {
		return err
	}

	return nil
}
