package service

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/pkg/log"
	"go.uber.org/zap"
)

type PostsService struct {
	repository post.Repository
	cache      post.Cache
}

func NewPostsService(repository post.Repository, cache post.Cache) *PostsService {
	return &PostsService{repository, cache}
}

func (p PostsService) List() ([]*post.Response, error) {
	posts, err := p.repository.List()
	if err != nil {
		return nil, err
	}

	response := post.ConvertEntitiesToResponses(posts)

	return response, nil
}

func (p PostsService) Get(id uuid.UUID) (*post.Response, error) {
	cachedPost, err := p.cache.Get(id)
	if err == nil {
		return post.ConvertEntityToResponse(cachedPost), nil
	}

	savedPost, err := p.repository.Get(id)
	if err != nil {
		return nil, err
	}

	go func() {
		if cacheErr := p.cache.Set(id, *savedPost); cacheErr != nil {
			log.Logger.Error("Error caching post: %v", zap.Error(cacheErr))
		}
	}()

	return post.ConvertEntityToResponse(*savedPost), nil
}

func (p PostsService) Create(params sqlc.CreatePostParams) (*post.Response, error) {
	savedPost, err := p.repository.Create(params)
	if err != nil {
		return nil, err
	}

	response := post.ConvertEntityToResponse(*savedPost)

	return response, nil
}
