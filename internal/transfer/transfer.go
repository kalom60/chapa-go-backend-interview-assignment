package transfer

import (
	"github.com/Chapa-Et/chapa-go"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/cache"
)

func New(webhooksecret string, transferRepo TransferRepo, chapaAPI chapa.API, redis cache.RedisCache) Handler {
	svc := NewService(transferRepo, chapaAPI, redis)
	return NewHandler(webhooksecret, svc)
}
