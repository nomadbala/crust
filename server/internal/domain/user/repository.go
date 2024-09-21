package user

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Repository interface {
	List() ([]sqlc.User, error)
	Get(uuid pgtype.UUID) (sqlc.User, error)
	Create(request SaltedRegistrationRequest) (sqlc.User, error)
}
