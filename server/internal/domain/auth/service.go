package auth

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/internal/domain/user"
)

type Service interface {
	SignUp(request RegistrationRequest) (*user.Response, error)
	SignIn(request LoginRequest) (*string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
}
