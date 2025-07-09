package transfer

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/cache"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/clients"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

var (
	ErrDuplicateTransfer = errors.New("duplicate transfer")
	ErrFailedInitiate    error
)

type Service struct {
	transferRepo TransferRepo
	chapa        clients.ChapaClient
	redis        cache.RedisCache
}

func NewService(transferRepo TransferRepo, chapa clients.ChapaClient, redis cache.RedisCache) *Service {
	return &Service{
		transferRepo: transferRepo,
		chapa:        chapa,
		redis:        redis,
	}
}

func (s *Service) InitiateTransfer(ctx context.Context, transfer clients.TransferRequest) (string, error) {
	ref := uuid.New()

	if _, err := s.redis.Get(ref.String()); err != nil {
		return "", ErrDuplicateTransfer
	}

	tr, err := createTransfer(transfer, ref.String())
	if err != nil {
		return "", err
	}

	err = s.redis.Set(ref.String(), tr, 15*time.Minute)
	if err != nil {
		return "", err
	}

	transfer.Reference = ref.String()
	resp, err := s.chapa.InitiateTransfer(transfer)
	if err != nil {
		return "", err
	}

	if resp.Status == "failed" {
		err := s.redis.Delete(ref.String())
		if err != nil {
			return "", err
		}

		ErrFailedInitiate = errors.New(resp.Message)
		return "", ErrFailedInitiate
	}

	err = s.transferRepo.CreateTransfer(ctx, tr)
	if err != nil {
		return "", err
	}

	return ref.String(), nil
}

func (s *Service) GetAllTransfers(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransfers[Transfer], error) {
	return s.transferRepo.GetAllTransfers(ctx, filter)
}

func (s *Service) VerifyTransfer(ref string) (*clients.VerifyResponse, error) {
	return s.chapa.VerifyTransfer(ref)
}
