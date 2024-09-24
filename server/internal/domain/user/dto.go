package user

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
)

type Response struct {
	ID                 uuid.UUID        `db:"id" json:"id"`
	Username           string           `db:"username" json:"username"`
	Email              string           `db:"email" json:"email"`
	FirstName          *pgtype.Text     `db:"first_name" json:"first_name"`
	LastName           *pgtype.Text     `db:"last_name" json:"last_name"`
	PhoneNumber        *pgtype.Text     `db:"phone_number" json:"phone_number"`
	DateOfBirth        *pgtype.Date     `db:"date_of_birth" json:"date_of_birth"`
	Gender             *sqlc.NullGender `db:"gender" json:"gender"`
	Bio                *pgtype.Text
	LanguagePreference *sqlc.LanguagePreference `db:"language_preference" json:"language_preference"`
}

type UserCredentials struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Password string    `db:"password" json:"password"`
	Salt     string    `db:"salt" json:"salt"`
}

func ConvertEntityToResponse(entity sqlc.User) *Response {
	return &Response{
		ID:       entity.ID,
		Username: entity.Username,
	}
}

func ConvertEntitiesToResponses(entities []sqlc.User) []*Response {
	responses := make([]*Response, len(entities))

	for i, entity := range entities {
		responses[i] = ConvertEntityToResponse(entity)
	}

	return responses
}
