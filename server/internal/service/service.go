package service

import (
	"github.com/nomadbala/crust/server/internal/cache"
	"github.com/nomadbala/crust/server/internal/domain/auth"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"github.com/nomadbala/crust/server/internal/repository"
)

type Configuration struct {
	SigningKey string
}

type Service struct {
	UsersService user.Service
	AuthService  auth.Service
	PostsService post.Service
}

func New(r *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		UsersService: NewUsersService(r.UsersRepository),
		AuthService:  NewAuthenticationService(r.UsersRepository),
		PostsService: NewPostsService(r.PostsRepository, cache.PostCache),
	}
}
