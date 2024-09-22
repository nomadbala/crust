package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5/pgtype"
	"net/mail"
)

type RegistrationRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func ParseRegistrationRequest(request RegistrationRequest) error {
	if request.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return err
	}

	return nil
}

type SaltedRegistrationRequest struct {
	RegistrationRequest
	Salt string
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId pgtype.UUID `json:"user_id"`
}
