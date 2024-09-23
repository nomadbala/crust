package service

import (
	"github.com/nomadbala/crust/server/internal/domain/auth"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"github.com/nomadbala/crust/server/internal/repository"
)

type Configuration func(s *Service) error

type Service struct {
	UsersService user.Service
	AuthService  auth.Service
	PostsService post.Service
}

func New(r *repository.Repository) *Service {
	return &Service{
		UsersService: NewUsersService(r.UsersRepository),
		AuthService:  NewAuthenticationService(r.UsersRepository),
		PostsService: NewPostsService(r.PostsRepository),
	}
}
