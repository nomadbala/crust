package cache

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/pkg/store"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

type VerificationCache struct {
	client       *store.RedisClient
	CACHE_PREFIX string
}

func NewVerificationCache(client *store.RedisClient) *VerificationCache {
	return &VerificationCache{client, "user_id_"}
}

func (c VerificationCache) Get(id uuid.UUID) (string, error) {
	data, err := c.client.Get(c.CACHE_PREFIX + id.String())
	if err != nil && data == "" {
		return "", ErrorCacheNotFound
	}

	var value string
	if err := msgpack.Unmarshal([]byte(data), &value); err != nil {
		return "", err
	}

	return value, nil
}

func (c VerificationCache) Set(id uuid.UUID, value string) error {
	packed, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(c.CACHE_PREFIX+id.String(), packed, 5*time.Minute)
}

func (c VerificationCache) Delete(id uuid.UUID) error {
	err := c.client.Del(c.CACHE_PREFIX + id.String())
	if err != nil {
		return err
	}

	return nil
}
