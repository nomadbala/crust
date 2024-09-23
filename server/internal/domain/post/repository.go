package post

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Repository interface {
	List() ([]sqlc.Post, error)
	Get(id pgtype.UUID) (sqlc.Post, error)
	GetPopular(params sqlc.GetPopularPostsParams) ([]sqlc.Post, error)
	Create(sqlc.CreatePostParams) (sqlc.Post, error)
}
