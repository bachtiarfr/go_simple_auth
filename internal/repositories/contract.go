package repositories

import "simple-auth/internal/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]entity.User, error)
}
