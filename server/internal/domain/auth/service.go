package auth

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/internal/domain/user"
)

type Service interface {
	SignUp(request RegistrationRequest) (*user.Response, error)
	SignIn(request LoginRequest) (string, error)
	ParseToken(accessToken string) (pgtype.UUID, error)
}
