package user

type Service interface {
	List() ([]*Response, error)
	//SendEmailVerification(id uuid.UUID) (bool, error)
}
