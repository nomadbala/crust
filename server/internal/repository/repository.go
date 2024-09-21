package repository

import (
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/domain/user"
)

type Configuration func(r *Repository) error

type Repository struct {
	queries         *sqlc.Queries
	UsersRepository user.Repository
}

func New(queries *sqlc.Queries) *Repository {
	return &Repository{
		UsersRepository: NewUsersRepository(queries),
	}
}
