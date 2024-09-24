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

//func (s *UsersService) SendEmailVerification(id uuid.UUID) (bool, error) {
//	receiver, err := s.repository.GetEmailById(id)
//	if err != nil {
//		return false, err
//	}
//
//	err = email.SendVerificationEmail(receiver)
//	if err != nil {
//		return false, err
//	}
//
//	return true, nil
//}
