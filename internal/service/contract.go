package services

import (
	"simple-auth/internal/entity"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationService interface {
	Authenticate(email, password string) (string, error)
}

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	ValidateToken(tokenString string) (jwt.Claims, error)
}
