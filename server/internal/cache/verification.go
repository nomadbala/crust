package cache

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/pkg/store"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

type VerificationCache struct {
	client *store.RedisClient
}

func NewVerificationCache(client *store.RedisClient) *VerificationCache {
	return &VerificationCache{client}
}

func (c VerificationCache) Get(id uuid.UUID) (string, error) {
	data, err := c.client.Get("user_id_" + id.String())
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

	return c.client.Set("user_id_"+id.String(), packed, 5*time.Minute)
}

func (c VerificationCache) Delete(id uuid.UUID) error {
	err := c.client.Del("user_id_" + id.String())
	if err != nil {
		return err
	}

	return nil
}
