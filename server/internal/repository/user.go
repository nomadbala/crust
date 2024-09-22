package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type UsersRepository struct {
	queries *sqlc.Queries
	ctx     context.Context
}

func NewUsersRepository(queries *sqlc.Queries, ctx context.Context) *UsersRepository {
	return &UsersRepository{queries, ctx}
}

func (u UsersRepository) List() ([]sqlc.User, error) {
	users, err := u.queries.ListUsers(u.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UsersRepository) Get(username string) (id pgtype.UUID, password, salt string, err error) {
	var getUserRow sqlc.GetUserRow
	getUserRow, err = u.queries.GetUser(u.ctx, username)
	if err != nil {
		return pgtype.UUID{}, "", "", err
	}

	return getUserRow.ID, getUserRow.PasswordHash, getUserRow.Salt, nil
}

func (u UsersRepository) GetById(uuid pgtype.UUID) (sqlc.User, error) {
	user, err := u.queries.GetUserById(u.ctx, uuid)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (u UsersRepository) Create(params sqlc.CreateUserParams) (sqlc.User, error) {
	user, err := u.queries.CreateUser(u.ctx, params)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
