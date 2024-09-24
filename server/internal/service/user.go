package service

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"github.com/nomadbala/crust/server/internal/domain/verification"
	"github.com/nomadbala/crust/server/pkg/email"
)

type UsersService struct {
	repository user.Repository
	cache      verification.Cache
}

func NewUsersService(repository user.Repository, cache verification.Cache) *UsersService {

	return &UsersService{repository, cache}
}

func (s *UsersService) List() ([]*user.Response, error) {
	users, err := s.repository.List()
	if err != nil {
		return nil, err
	}

	responses := user.ConvertEntitiesToResponses(users)

	return responses, nil
}

func (s *UsersService) SendEmailVerification(id uuid.UUID) (bool, error) {
	receiver, err := s.repository.GetEmailById(id)
	if err != nil {
		return false, err
	}

	code := email.GenerateVerificationCode()
	err = email.SendVerificationEmail(receiver, code)
	if err != nil {
		return false, err
	}

	err = s.cache.Set(id, code)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *UsersService) VerifyEmail(id uuid.UUID, code string) (bool, error) {
	cachedCode, err := s.cache.Get(id)
	if err != nil {
		return false, err
	}

	if cachedCode != code {
		return false, nil
	}

	if err := s.repository.VerifyEmail(id); err != nil {
		return false, err
	}

	if err := s.cache.Delete(id); err != nil {
		return false, err
	}

	return true, nil
}
