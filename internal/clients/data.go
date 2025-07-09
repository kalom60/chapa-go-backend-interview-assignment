package clients

import "time"

type TransferRequest struct {
	Charge        float64 `json:"charge,omitempty"`
	AccountName   string  `json:"account_name,omitempty"`
	AccountNumber string  `json:"account_number" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Currency      string  `json:"currency,omitempty"`
	Reference     string  `json:"reference,omitempty"`
	BankCode      string  `json:"bank_code" binding:"required"`
}

type TransferResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    string `json:"data"`
}

type VerifyResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		AccountName     string    `json:"account_name"`
		AccountNumber   string    `json:"account_number"`
		Currency        string    `json:"currency"`
		Amount          float64   `json:"amount"`
		Charge          float64   `json:"charge"`
		TransferMethod  string    `json:"transfer_method"`
		ChapaTransferID string    `json:"chapa_transfer_id"`
		BankCode        string    `json:"bank_code"`
		BankName        string    `json:"bank_name"`
		Status          string    `json:"status"`
		TxRef           string    `json:"tx_ref"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	} `json:"data"`
}
