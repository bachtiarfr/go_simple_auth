package repositories

import "simple-auth/internal/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]entity.ListUser, error)
	CreateNewUser(user *entity.User) (entity.User, error)
}
