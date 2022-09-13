package login

type LoginCore struct {
	ID       int
	Email    string
	Password string
}

type UsecaseInterface interface {
	LoginAuthorized(email, password string) (string, error)
}

type DataInterface interface {
	LoginUser(email, password string) (LoginCore, error)
}