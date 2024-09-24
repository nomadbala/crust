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

func (p PostCache) Get(id uuid.UUID) (sqlc.Post, error) {
	data, err := p.client.Get(id.String())
	if err != nil {
		return sqlc.Post{}, err
	}

	var post sqlc.Post
	if err := msgpack.Unmarshal([]byte(data), &post); err != nil {
		return sqlc.Post{}, err
	}

	return post, nil
}

func (p PostCache) Set(key uuid.UUID, value sqlc.Post) error {
	packed, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}

	return p.client.Set(key.String(), packed, 15*time.Minute)
}

func (p PostCache) Del(key uuid.UUID) error {
	err := p.client.Del(key.String())
	if err != nil {
		return err
	}

	return nil
}
