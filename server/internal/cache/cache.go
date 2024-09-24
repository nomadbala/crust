package cache

import (
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/pkg/store"
)

type Configuration func(c *Cache) error

type Cache struct {
	PostCache post.Cache
}

func New(client *store.RedisClient) *Cache {
	return &Cache{
		PostCache: NewPostCache(client),
	}
}
