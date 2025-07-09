package transaction

import (
	"time"

	"github.com/Chapa-Et/chapa-go"
)

type Transaction struct {
	Status            string    `json:"status"`
	RefID             string    `json:"ref_id"`
	Type              string    `json:"type"`
	CreatedAt         time.Time `json:"created_at"`
	Currency          string    `json:"currency"`
	Amount            float64   `json:"amount"`
	Charge            float64   `json:"charge"`
	TransID           string    `json:"trans_id"`
	PaymentMethod     string    `json:"payment_method"`
	CustomerID        string    `json:"customer_id"`
	CustomerFirstname string    `json:"customer_firstname"`
	CustomerLastname  string    `json:"customer_lastname"`
	CustomerMobile    string    `json:"customer_mobile"`
}

func createTransaction(tx *chapa.PaymentRequest, ref string) (Transaction, error) {
	return Transaction{
		Status:    "proccessing",
		RefID:     ref,
		CreatedAt: time.Now(),
		Currency:  tx.Currency,
		Amount:    tx.Amount.InexactFloat64(),
	}, nil
}
