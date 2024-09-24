package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type PostsRepository struct {
	queries *sqlc.Queries
	ctx     context.Context
}

func NewPostsRepository(queries *sqlc.Queries, ctx context.Context) PostsRepository {
	return PostsRepository{queries, ctx}
}

func (p PostsRepository) List() ([]sqlc.Post, error) {
	posts, err := p.queries.ListPosts(p.ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p PostsRepository) Get(id uuid.UUID) (sqlc.Post, error) {
	post, err := p.queries.GetPostById(p.ctx, id)
	if err != nil {
		return sqlc.Post{}, err
	}

	return post, nil
}

func (p PostsRepository) GetPopular(params sqlc.GetPopularPostsParams) ([]sqlc.Post, error) {
	posts, err := p.queries.GetPopularPosts(p.ctx, params)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p PostsRepository) Create(params sqlc.CreatePostParams) (sqlc.Post, error) {
	post, err := p.queries.CreatePost(p.ctx, params)
	if err != nil {
		return sqlc.Post{}, err
	}

	return post, nil
}
