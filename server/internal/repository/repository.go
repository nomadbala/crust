package repository

import (
	"context"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/internal/domain/user"
)

type Configuration func(r *Repository) error

type Repository struct {
	queries         *sqlc.Queries
	UsersRepository user.Repository
	PostsRepository post.Repository
}

func New(queries *sqlc.Queries, ctx context.Context) *Repository {
	return &Repository{
		UsersRepository: NewUsersRepository(queries, ctx),
		PostsRepository: NewPostsRepository(queries, ctx),
	}
}
