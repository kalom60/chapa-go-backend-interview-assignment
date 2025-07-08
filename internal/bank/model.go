package bank

import "time"

type Bank struct {
	BankID        int       `json:"bank_id"`
	Slug          string    `json:"slug"`
	Swift         string    `json:"swift"`
	Name          string    `json:"name"`
	AcctLength    int       `json:"acct_length"`
	CountryID     int       `json:"country_id"`
	IsMobilemoney int       `json:"is_mobilemoney"`
	IsActive      int       `json:"is_active"`
	Is_Rtgs       int       `json:"is_rtgs"`
	Active        int       `json:"active"`
	Is24hrs       int       `json:"is_24hrs"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Currency      string    `json:"currency"`
}
