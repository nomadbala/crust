package service

import (
	"github.com/nomadbala/crust/server/internal/cache"
	"github.com/nomadbala/crust/server/internal/config"
	"github.com/nomadbala/crust/server/internal/domain/auth"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"github.com/nomadbala/crust/server/internal/repository"
)

type Configuration struct{}

type Service struct {
	UsersService user.Service
	AuthService  auth.Service
	PostsService post.Service
}

func New(r *repository.Repository, cache *cache.Cache, cfg config.Config) *Service {
	return &Service{
		UsersService: NewUsersService(r.UsersRepository, cache.VerificationCache),
		AuthService:  NewAuthenticationService(r.UsersRepository, cfg.Token),
		PostsService: NewPostsService(r.PostsRepository, cache.PostCache),
	}
}
