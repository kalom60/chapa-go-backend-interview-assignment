package transaction

import (
	"github.com/Chapa-Et/chapa-go"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/cache"
)

func New(webhooksecret string, txRepo TransactionRepo, chapaAPI chapa.API, redis cache.RedisCache) Handler {
	svc := NewService(txRepo, chapaAPI, redis)
	return NewHandler(webhooksecret, svc)
}
