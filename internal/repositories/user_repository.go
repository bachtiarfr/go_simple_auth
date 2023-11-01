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

func (r *userRepository) FindAll() ([]entity.User, error) {
	query := "SELECT id, fullname, email, age, phone_number FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Age, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
