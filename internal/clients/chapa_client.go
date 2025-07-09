package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	acceptPaymentV1APIURL  = "https://api.chapa.co/v1/transaction/initialize"
	transferV1APIURL       = "/transfers"
	verifyTransferV1APIURL = "/transaction/verify"
	transactionsV1APIURL   = "https://api.chapa.co/v1/transactions"
	banksV1APIURL          = "https://api.chapa.co/v1/banks"
	bulkTransferAPIURL     = "https://api.chapa.co/v1/bulk-transfers"
)

type ChapaClient interface {
	// PaymentRequest(request *PaymentRequest) (*PaymentResponse, error)
	InitiateTransfer(payload TransferRequest) (*TransferResponse, error)
	VerifyTransfer(ref string) (*VerifyResponse, error)
	// GetTransactions() (*TransactionsResponse, error)
	// GetBanks() (*BanksResponse, error)
}

type chapaClient struct {
	BaseURL   string
	SecretKey string
	Client    *http.Client
}

func NewChapaClient(baseUrl, secretkey string) ChapaClient {
	return &chapaClient{
		BaseURL:   baseUrl,
		SecretKey: secretkey,
		Client:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (cc *chapaClient) InitiatePayment(payload map[string]any) ([]byte, error) {
	return cc.post("", payload)
}

func (cc *chapaClient) VerifyPayment(txRef string) ([]byte, error) {
	url := fmt.Sprintf("%s/transaction/verify/%s", cc.BaseURL, txRef)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+cc.SecretKey)
	resp, err := cc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (cc *chapaClient) InitiateTransfer(payload TransferRequest) (*TransferResponse, error) {
	url := fmt.Sprintf("%s%s", cc.BaseURL, transferV1APIURL)

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+cc.SecretKey)
	req.Header.Set("Context-Type", "application/json")

	resp, err := cc.Client.Do(req)
	if err != nil {
		log.Printf("error %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading resposne body %v", err)
		return nil, err
	}

	var response TransferResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Printf("error while unmarshaling  response %v", err)
		return nil, err
	}
	return &response, nil
}

func (cc *chapaClient) VerifyTransfer(transferID string) (*VerifyResponse, error) {
	url := fmt.Sprintf("%s%s/%s", cc.BaseURL, verifyTransferV1APIURL, transferID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+cc.SecretKey)
	resp, err := cc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var verifyResponse VerifyResponse

	err = json.Unmarshal(body, &verifyResponse)
	if err != nil {
		return nil, err
	}

	return &verifyResponse, nil
}

func (cc *chapaClient) post(endpoint string, payload map[string]any) ([]byte, error) {
	url := fmt.Sprintf("%s%s", cc.BaseURL, endpoint)

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+cc.SecretKey)
	req.Header.Set("Context-Type", "application/json")

	resp, err := cc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Chapa API error: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
