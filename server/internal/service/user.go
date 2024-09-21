package service

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repository user.Repository
}

func NewUsersService(repository user.Repository) *UsersService {
	return &UsersService{repository: repository}
}

func (s *UsersService) List() ([]*user.Response, error) {
	users, err := s.repository.List()
	if err != nil {
		return nil, err
	}

	responses := user.ConvertEntitiesToResponses(users)

	return responses, nil
}

func (s *UsersService) SignUp(request user.RegistrationRequest) (*user.Response, error) {
	var registrationRequestDAO user.SaltedRegistrationRequest

	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashedPassword)

	registrationRequestDAO = user.SaltedRegistrationRequest{
		RegistrationRequest: request,
		Salt:                salt,
	}

	savedUser, err := s.repository.Create(registrationRequestDAO)
	if err != nil {
		return nil, err
	}

	return user.ConvertEntityToResponse(savedUser), nil
}

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(salt), nil
}
