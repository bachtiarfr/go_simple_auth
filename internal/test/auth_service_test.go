package services

import (
	"database/sql"
	"simple-auth/internal/repositories"
	services "simple-auth/internal/service"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAuthenticationService_Authenticate(t *testing.T) {
	// Buat database mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Inisialisasi repository dan service
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthenticationService(userRepository)

	// Kasus uji 1: Autentikasi berhasil
	mock.ExpectQuery("SELECT id, email, password FROM users").WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).
			AddRow(1, "test@example.com", "$2a$10$aUfnbFX8QEX.Rg3Dz/XegOhnh9hliMxdB3FRAcHg35af2YAyBSN9a"))

	token, err := authService.Authenticate("test@example.com", "password")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got an empty string")
	}

	// Kasus uji 2: Autentikasi gagal
	mock.ExpectQuery("SELECT id, email, password FROM users").WithArgs("test@example.com").
		WillReturnError(sql.ErrNoRows)

	token, err = authService.Authenticate("test@example.com", "wrongpassword")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
	if token != "" {
		t.Error("Expected an empty token, got a token")
	}

	// Verifikasi bahwa semua harapan panggilan database telah dipenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
