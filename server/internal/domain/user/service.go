package user

import "github.com/google/uuid"

type Service interface {
	List() ([]*Response, error)
	SendEmailVerification(id uuid.UUID) (bool, error)
	VerifyEmail(id uuid.UUID, code string) (bool, error)
}
