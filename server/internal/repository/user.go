package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type UsersRepository struct {
	queries *sqlc.Queries
	ctx     context.Context
}

func NewUsersRepository(queries *sqlc.Queries, ctx context.Context) *UsersRepository {
	return &UsersRepository{queries, ctx}
}

func (r UsersRepository) List() ([]sqlc.User, error) {
	users, err := r.queries.ListUsers(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r UsersRepository) Get(username string) (id uuid.UUID, password, salt string, err error) {
	var getUserRow sqlc.GetUserRow
	getUserRow, err = r.queries.GetUser(r.ctx, username)
	if err != nil {
		return uuid.Nil, "", "", err
	}

	return getUserRow.ID, getUserRow.PasswordHash, getUserRow.Salt, nil
}

func (r UsersRepository) GetById(uuid uuid.UUID) (sqlc.User, error) {
	user, err := r.queries.GetUserById(r.ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

//func (r UsersRepository) GetEmailById(uuid uuid.UUID) (string, error) {
//	email, err := r.queries.GetEmailById(r.ctx, uuid)
//	if err != nil {
//		return "", err
//	}
//
//	return email, nil
//}

func (r UsersRepository) Create(params sqlc.CreateUserParams) (sqlc.User, error) {
	user, err := r.queries.CreateUser(r.ctx, params)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
