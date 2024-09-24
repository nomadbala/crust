package cache

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/pkg/store"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

type PostCache struct {
	client *store.RedisClient
}

func NewPostCache(client *store.RedisClient) *PostCache {
	return &PostCache{client}
}

func (c PostCache) Get(id uuid.UUID) (sqlc.Post, error) {
	data, err := c.client.Get("post_id_" + id.String())
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

	return c.client.Set("post_id_"+key.String(), packed, 15*time.Minute)
}

func (c PostCache) Del(key uuid.UUID) error {
	err := c.client.Del("post_id_" + key.String())
	if err != nil {
		return err
	}

	return nil
}
