package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/config"
	"github.com/nomadbala/crust/server/internal/domain/auth"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticationService struct {
	repository user.Repository
	cfg        config.Token
}

var (
	ErrorInvalidUsernameOrPassword = errors.New("unable to sign in. Username or password are invalid. Try again")
	ErrorTokenGenerationFailed     = errors.New("failed to generate a valid token. Please try again or contact support")
)

func NewAuthenticationService(repository user.Repository, cfg config.Token) *AuthenticationService {
	return &AuthenticationService{repository, cfg}
}

func (s *AuthenticationService) SignUp(request auth.RegistrationRequest) (*user.Response, error) {
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
		PasswordHash: hashedPassword,
		Salt:         salt,
		Email:        request.Email,
	}

	savedUser, err := s.repository.Create(params)
	if err != nil {
		return nil, err
	}

	return user.ConvertEntityToResponse(*savedUser), nil
}

func (s *AuthenticationService) SignIn(request auth.LoginRequest) (*string, error) {
	credentials, err := s.repository.Get(request.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(request.Password+credentials.Salt)); err != nil {
		return nil, ErrorInvalidUsernameOrPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.cfg.Expires).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: credentials.ID,
	})

	accessToken, err := token.SignedString([]byte(s.cfg.SigningKey))
	if err != nil || accessToken == "" {
		return nil, ErrorTokenGenerationFailed
	}

	return &accessToken, nil
}

func (s *AuthenticationService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.SigningKey), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*auth.TokenClaims)

	if !ok || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
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
