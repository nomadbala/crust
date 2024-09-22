package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/domain/auth"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var (
	signinKey = os.Getenv("JWT_SIGNIN_KEY")
)

type AuthenticationService struct {
	repository user.Repository
}

func NewAuthenticationService(repository user.Repository) *AuthenticationService {
	return &AuthenticationService{repository}
}

func (a AuthenticationService) SignUp(request auth.RegistrationRequest) (*user.Response, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(request.Password, salt)
	if err != nil {
		return nil, err
	}

	params := sqlc.CreateUserParams{
		Username:     request.Username,
		PasswordHash: string(hashedPassword),
		Salt:         salt,
	}

	savedUser, err := a.repository.Create(params)
	if err != nil {
		return nil, err
	}

	return user.ConvertEntityToResponse(savedUser), nil
}

func (a AuthenticationService) SignIn(request auth.LoginRequest) (string, error) {
	id, passwordHashFromDB, salt, err := a.repository.Get(request.Username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHashFromDB), []byte(request.Password+salt)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	tokenString, err := token.SignedString([]byte(signinKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthenticationService) ParseToken(accessToken string) (pgtype.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signin method")
		}

		return []byte(signinKey), nil
	})

	if err != nil {
		return pgtype.UUID{}, err
	}

	claims, ok := token.Claims.(*auth.TokenClaims)

	if !ok || !token.Valid {
		return pgtype.UUID{}, errors.New("invalid token")
	}

	return claims.UserId, nil
}

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(salt), nil
}

func hashPassword(password, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
