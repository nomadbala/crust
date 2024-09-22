package user

type Service interface {
	List() ([]*Response, error)
}
