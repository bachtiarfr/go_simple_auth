package repositories

import (
	"database/sql"
	"simple-auth/internal/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = $1"
	var user entity.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Pengguna tidak ditemukan
		}
		return nil, err
	}

	return &user, nil
}
