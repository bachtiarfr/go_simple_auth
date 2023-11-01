package services

type AuthenticationService interface {
	Authenticate(email, password string) (string, error)
}
