package transfer

import (
	"strconv"
	"time"

	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/clients"
)

type Transfer struct {
	AccountName    string    `json:"account_name"`
	AccountNumber  string    `json:"account_number"`
	Currency       string    `json:"currency"`
	Amount         float64   `json:"amount"`
	Charge         float64   `json:"charge"`
	TransferType   string    `json:"transfer_type"`
	ChapaReference string    `json:"chapa_reference"`
	BankCode       int       `json:"bank_code"`
	BankName       string    `json:"bank_name"`
	BankReference  string    `json:"bank_reference"`
	Status         string    `json:"status"`
	Reference      string    `json:"reference"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func createTransfer(transfer clients.TransferRequest, ref string) (Transfer, error) {
	bankCode, err := strconv.Atoi(transfer.BankCode)
	if err != nil {
		return Transfer{}, err
	}

	return Transfer{
		AccountName:   transfer.AccountName,
		AccountNumber: transfer.AccountNumber,
		Currency:      transfer.Currency,
		Amount:        float64(transfer.Amount),
		BankCode:      bankCode,
		Status:        "proccessing",
		Reference:     ref,
		CreatedAt:     time.Now(),
	}, nil
}
