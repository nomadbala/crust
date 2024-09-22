package user

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Repository interface {
	List() ([]sqlc.User, error)
	GetById(uuid pgtype.UUID) (sqlc.User, error)
	Get(username string) (id pgtype.UUID, password, salt string, err error)
	Create(params sqlc.CreateUserParams) (sqlc.User, error)
}
