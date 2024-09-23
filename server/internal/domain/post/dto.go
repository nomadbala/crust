package post

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type CreatePostRequest struct {
	UserId  pgtype.UUID `json:"user_id"`
	Content string      `json:"content"`
}

type Response struct {
	Id        pgtype.UUID      `json:"id"`
	UserId    pgtype.UUID      `json:"user_id"`
	Content   string           `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func ConvertEntityToResponse(entity sqlc.Post) *Response {
	return &Response{
		Id:        entity.ID,
		UserId:    entity.UserID,
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
	}
}

func ConvertEntitiesToResponses(entities []sqlc.Post) []*Response {
	responses := make([]*Response, len(entities))

	for i, entity := range entities {
		responses[i] = ConvertEntityToResponse(entity)
	}

	return responses
}
