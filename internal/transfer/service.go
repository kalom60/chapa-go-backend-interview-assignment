package transfer

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
	ErrDuplicateTransfer = errors.New("duplicate transfer")
	ErrFailedInitiate    error
)

type Service struct {
	transferRepo TransferRepo
	chapaAPI     chapa.API
	redis        cache.RedisCache
}

func NewService(transferRepo TransferRepo, chapaAPI chapa.API, redis cache.RedisCache) *Service {
	return &Service{
		transferRepo: transferRepo,
		chapaAPI:     chapaAPI,
		redis:        redis,
	}
}

func (s *Service) InitiateTransfer(ctx context.Context, transfer *chapa.BankTransfer) (string, error) {
	ref := uuid.New()

	val, err := s.redis.Get(ref.String())
	if err != nil {
		if err != redis.Nil {
			return "", err
		}
	}

	if val != nil {
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
	resp, err := s.chapaAPI.TransferToBank(transfer)
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

	return resp.Data, nil
}

func (s *Service) GetAllTransfers(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseTransfers[Transfer], error) {
	return s.transferRepo.GetAllTransfers(ctx, filter)
}

func (s *Service) VerifyTransfer(ref string) (*chapa.VerifyResponse, error) {
	return s.chapaAPI.Verify(ref)
}

func (s *Service) HandleWebhook(ctx context.Context, transfer Transfer) error {
	if _, err := s.redis.Get(transfer.Reference); err != nil {
		if err == redis.Nil {
			_, err := s.transferRepo.GetTransferByRef(ctx, transfer.Reference)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return err
				}
				return err
			}
		}
		return err
	}

	err := s.transferRepo.UpdateTransfer(ctx, transfer)
	if err != nil {
		return err
	}

	err = s.redis.Delete(transfer.Reference)
	if err != nil {
		return err
	}

	return nil
}
