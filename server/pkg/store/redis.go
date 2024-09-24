package store

import (
	"context"
	"github.com/nomadbala/crust/server/pkg/log"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	ctx    context.Context
	client *redis.Client
}

func (r *RedisClient) New(ctx context.Context, url string) {
	options, err := redis.ParseURL(url)
	if err != nil {
		log.Logger.Error("error occurred while parsing redis url")
		return
	}

	r.ctx = ctx
	r.client = redis.NewClient(options)
}

func (r *RedisClient) Close() {
	err := r.client.Close()
	if err != nil {
		return
	}
}

func (r *RedisClient) Get(key string) (string, error) {
	data, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}

	return data, nil
}

func (r *RedisClient) Set(key string, value []byte, expiration time.Duration) error {
	if err := r.client.Set(r.ctx, key, value, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Del(key string) error {
	if err := r.client.Del(r.ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
