package verification

import "github.com/google/uuid"

type Cache interface {
	Get(id uuid.UUID) (string, error)
	Set(id uuid.UUID, value string) error
	Delete(id uuid.UUID) error
}
