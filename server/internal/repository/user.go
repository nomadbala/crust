package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/domain/user"
)

type UsersRepository struct {
	queries *sqlc.Queries
}

var ctx = context.Background()

func (u UsersRepository) List() ([]sqlc.User, error) {
	users, err := u.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UsersRepository) Get(uuid pgtype.UUID) (sqlc.User, error) {
	user, err := u.queries.GetUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (u UsersRepository) Create(request user.SaltedRegistrationRequest) (sqlc.User, error) {
	user, err := u.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Username:     request.Username,
		PasswordHash: request.Password,
		Salt:         request.Salt,
	})
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func NewUsersRepository(db *sqlc.Queries) *UsersRepository {
	return &UsersRepository{queries: db}
}
