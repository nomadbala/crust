package user

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Repository interface {
	List() ([]sqlc.User, error)
	GetById(uuid uuid.UUID) (*sqlc.User, error)
	GetEmailById(uuid uuid.UUID) (*string, error)
	Get(username string) (*UserCredentials, error)
	Create(params sqlc.CreateUserParams) (*sqlc.User, error)
	VerifyEmail(id uuid.UUID) error
}
