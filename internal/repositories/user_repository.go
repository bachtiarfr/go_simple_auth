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
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll() ([]entity.ListUser, error) {
	query := "SELECT id, fullname, email, age, phone_number FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.ListUser
	for rows.Next() {
		var user entity.ListUser
		err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Age, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) CreateNewUser(user *entity.User) (entity.User, error) {
	query := "INSERT INTO users (fullname, email, age, password, phone_number) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	// Eksekusi query dan ambil ID pengguna yang baru saja dibuat
	err := r.db.QueryRow(query, user.Fullname, user.Email, user.Age, user.Password, user.PhoneNumber).Scan(&user.ID)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{}, err
}
