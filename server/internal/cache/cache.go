package cache

import (
	"errors"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/internal/domain/verification"
	"github.com/nomadbala/crust/server/pkg/store"
)

type Configuration func(c *Cache) error

type Cache struct {
	PostCache         post.Cache
	VerificationCache verification.Cache
}

var ErrorCacheNotFound = errors.New("cache miss: requested item not found in cache")

func New(client *store.RedisClient) *Cache {
	return &Cache{
		PostCache:         NewPostCache(client),
		VerificationCache: NewVerificationCache(client),
	}
}
