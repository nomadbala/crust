package post

import (
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Service interface {
	List() ([]*Response, error)
	Get(id uuid.UUID) (*Response, error)
	Create(params sqlc.CreatePostParams) (*Response, error)
}
