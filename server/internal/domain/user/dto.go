package user

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"net/mail"
)

type RegistrationRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func ParseRegistrationRequest(request RegistrationRequest) error {
	_, err := mail.ParseAddress(request.Email)

	if err != nil {
		return err
	}
}

type SaltedRegistrationRequest struct {
	RegistrationRequest
	Salt string
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	ID                 pgtype.UUID      `db:"id" json:"id"`
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
