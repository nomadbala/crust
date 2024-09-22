package service

import (
	"github.com/nomadbala/crust/server/internal/domain/user"
)

type UsersService struct {
	repository user.Repository
}

func NewUsersService(repository user.Repository) *UsersService {

	return &UsersService{repository}
}

func (s *UsersService) List() ([]*user.Response, error) {
	users, err := s.repository.List()
	if err != nil {
		return nil, err
	}

	responses := user.ConvertEntitiesToResponses(users)

	return responses, nil
}
