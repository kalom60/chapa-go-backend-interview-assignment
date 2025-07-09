package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedis(redisAddr, redisPassword string) (RedisCache, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &redisCache{
		client: client,
		ctx:    ctx,
	}, nil
}

func (r *redisCache) Set(key string, value any, ttl time.Duration) error {
	return r.client.Set(r.ctx, key, value, ttl).Err()
}

func (r *redisCache) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *redisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
