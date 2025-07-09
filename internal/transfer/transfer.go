package transfer

import (
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/cache"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/clients"
)

func New(webhooksecret string, transferRepo TransferRepo, chapaClient clients.ChapaClient, redis cache.RedisCache) Handler {
	svc := NewService(transferRepo, chapaClient, redis)
	return NewHandler(webhooksecret, svc)
}
