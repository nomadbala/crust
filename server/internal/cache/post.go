package cache

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/pkg/store"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

type PostCache struct {
	client       *store.RedisClient
	CACHE_PREFIX string
}

func NewPostCache(client *store.RedisClient) *PostCache {
	return &PostCache{client, "post_id_"}
}

func (c PostCache) Get(id uuid.UUID) (sqlc.Post, error) {
	data, err := c.client.Get(c.CACHE_PREFIX + id.String())
	if err != nil {
		return sqlc.Post{}, err
	}

	var post sqlc.Post
	if err := msgpack.Unmarshal([]byte(data), &post); err != nil {
		return sqlc.Post{}, err
	}

	return post, nil
}

func (c PostCache) Set(key uuid.UUID, value sqlc.Post) error {
	packed, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(c.CACHE_PREFIX+key.String(), packed, 15*time.Minute)
}

func (c PostCache) Del(key uuid.UUID) error {
	err := c.client.Del(c.CACHE_PREFIX + key.String())
	if err != nil {
		return err
	}

	return nil
}
