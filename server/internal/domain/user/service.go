package user

type Service interface {
	List() ([]*Response, error)
	SignUp(request RegistrationRequest) (*Response, error)
}
