package post

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Cache interface {
	Get(id uuid.UUID) (sqlc.Post, error)
	Set(key uuid.UUID, value sqlc.Post) error
	Del(key uuid.UUID) error
}
