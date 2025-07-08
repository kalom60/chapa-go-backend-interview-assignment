package bank

import (
	"context"

	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

type Service struct {
	bankRepo BankRepo
}

func NewService(bankRepo BankRepo) *Service {
	return &Service{
		bankRepo: bankRepo,
	}
}

func (s *Service) GetAllBanks(ctx context.Context, filter utils.Pagination) (utils.PaginatedResponseBanks[Bank], error) {
	return s.bankRepo.GetAllBanks(ctx, filter)
}

func (s *Service) GetBankByBankID(ctx context.Context, bankID int) (Bank, error) {
	return s.bankRepo.GetBankByBankID(ctx, bankID)
}
