package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
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
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, key, data, ttl).Err()
}

func (r *redisCache) Get(key string) (any, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	return val, nil
}

func (r *redisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
